package config

import (
	"reflect"
	"strconv"

	ini "github.com/go-ini/ini"
)

type server struct {
	Id      string `ini:"id"`
	Port    int    `ini:"port"`
	TlsPort int    `ini:"tls_port"`
}

type log struct {
	Compress   bool   `ini:"compress"`
	LogPath    string `ini:"path"`
	MaxAge     int    `ini:"max_age"`
	MaxBackups int    `ini:"maxbackups"`
	MaxSize    int    `ini:"maxsize"`
}

type debug struct {
	ProfileOn   bool   `ini:"profile_on"`
	ProfilePort string `ini:"profile_port"`
}

type iniConfig struct {
	Server server `ini:"server"`
	Log    log    `ini:"log"`
	Dbg    debug  `ini:"debug"`
}

var thisConfig = iniConfig{}

func InitFromFile(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}

	err = cfg.MapTo(&thisConfig)
	if err != nil {
		return err
	}

	createIndex()
	return nil
}

func GetSectionKey(name string) (interface{}, bool) {
	v, ok := index[name]
	return v, ok
}

func GetInt(name string) (int, bool) {
	i, ok := GetSectionKey(name)
	if !ok {
		return 0, false
	}

	if v, ok := i.(int); ok {
		return v, true
	}

	// maybe it is a digital string
	s, ok := i.(string)
	if !ok {
		return 0, false
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return n, true
}

func GetString(name string) (string, bool) {
	i, ok := GetSectionKey(name)
	if !ok {
		return "", false
	}

	s, ok := i.(string)
	if !ok {
		return "", false
	}
	return s, true
}

func GetBool(name string) (bool, bool) {
	i, ok := GetSectionKey(name)
	if !ok {
		return false, false
	}

	b, ok := i.(bool)
	if !ok {
		return false, false
	}
	return b, true
}

var index = make(map[string]interface{}, 100)

func createIndex() {
	v := reflect.ValueOf(thisConfig)
	t := reflect.TypeOf(thisConfig)
	cnt := v.NumField()
	for i := 0; i < cnt; i++ {
		fieldVal := v.Field(i)
		if fieldVal.Kind() != reflect.Struct {
			continue
		}

		// it is a struct kind field, go on to get tag
		fieldStructTyp := t.Field(i)
		tag := fieldStructTyp.Tag.Get("ini")
		if tag == "" {
			continue // no ini tag, ignore it
		}

		// append Field Recursively
		appendField(tag, fieldVal)
	}
}

func appendField(parentTag string, v reflect.Value) {
	cnt := v.NumField()
	for i := 0; i < cnt; i++ {
		fieldVal := v.Field(i)
		fieldTyp := v.Type()
		fieldStructTyp := fieldTyp.Field(i)
		tag := fieldStructTyp.Tag.Get("ini")
		if tag == "" {
			continue
		}
		if fieldVal.Kind() != reflect.Struct {
			// leaf field, add to map
			index[parentTag+"."+tag] = fieldVal.Interface()
		} else {
			// recursive call appendField
			appendField(parentTag+"."+tag, fieldVal)
		}
	}
}

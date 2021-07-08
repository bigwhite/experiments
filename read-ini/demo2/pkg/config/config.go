package config

import (
	"reflect"
	"strconv"
	"strings"

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

	return cfg.MapTo(&thisConfig)
}

func GetSectionKey(name string) (interface{}, bool) {
	keys := strings.Split(name, ".")
	lastKey := keys[len(keys)-1]
	v := reflect.ValueOf(thisConfig)
	t := reflect.TypeOf(thisConfig)

	found := false
	for _, key := range keys {
		cnt := v.NumField()

		for i := 0; i < cnt; i++ {
			field := t.Field(i)
			if field.Tag.Get("ini") == key {
				t = field.Type
				v = v.Field(i)
				if key == lastKey {
					found = true
				}
				break
			}
		}
	}

	if found {
		return v.Interface(), true
	}
	return nil, false
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

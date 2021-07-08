package config

import (
	ini "github.com/go-ini/ini"
)

type Server struct {
	Id      string `ini:""`
	Port    int    `ini:"port"`
	TlsPort int    `ini:"tls_port"`
}

type Log struct {
	Compress   bool   `ini:"compress"`
	LogPath    string `ini:"path"`
	MaxAge     int    `ini:"max_age"`
	MaxBackups int    `ini:"maxbackups"`
	MaxSize    int    `ini:"maxsize"`
}

type Debug struct {
	ProfileOn   bool   `ini:"profile_on"`
	ProfilePort string `ini:"profile_port"`
}

type IniConfig struct {
	Server `ini:"server"`
	Log    `ini:"log"`
	Debug  `ini:"debug"`
}

var Config = &IniConfig{}

func InitFromFile(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}

	return cfg.MapTo(Config)
}

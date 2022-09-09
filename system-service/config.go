package main

import (
	"errors"
	"fmt"
	"os"

	ini "github.com/go-ini/ini"
	"github.com/kardianos/service"
)

func getServiceConfig(subCmd string) (*service.Config, error) {
	c := service.Config{
		Name:             "myApp",
		DisplayName:      "Go Daemon Service Demo",
		Description:      "This is a Go daemon service demo",
		Executable:       "/usr/local/bin/myapp",
		Dependencies:     []string{"After=network.target syslog.target"},
		WorkingDirectory: "",
		Option: service.KeyValue{
			"Restart": "always", // Restart=always
		},
	}

	switch subCmd {
	case "install":
		installCommand.Parse(os.Args[2:])
		if user == "" {
			fmt.Printf("error: user should be provided when install service\n")
			return nil, errors.New("invalid user")
		}
		if workingdir == "" {
			fmt.Printf("error: workingdir should be provided when install service\n")
			return nil, errors.New("invalid workingdir")
		}
		c.UserName = user
		c.WorkingDirectory = workingdir

		// arguments
		// ExecStart=/usr/local/bin/myapp "run" "-config" "/etc/myapp/config.ini"
		c.Arguments = append(c.Arguments, "run", "-config", config)
	case "run":
		runCommand.Parse(os.Args[2:]) // parse config
	}

	return &c, nil
}

type Server struct {
	Addr string `ini:"addr"`
}

type IniConfig struct {
	Server Server `ini:"server"`
}

func loadConfigFromIni(path string) (*IniConfig, error) {
	var c IniConfig
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	err = cfg.MapTo(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kardianos/service"
)

var (
	installCommand = flag.NewFlagSet("install", flag.ExitOnError)
	runCommand     = flag.NewFlagSet("run", flag.ExitOnError)
	user           string
	workingdir     string
	config         string
)

const (
	defaultConfig = "/etc/myapp/config.ini"
)

func usage() {
	s := `
USAGE:
   myapp command [command options] 


COMMANDS:
     install               install service
     uninstall             uninstall service
     start                 start service
     stop                  stop service
	 run                   run service

OPTIONS:
     -config string
       	config file of the service (default "/etc/myapp/config.ini")
     -user string
       	user account to run the service
     -workingdir string
    	working directory of the service`

	fmt.Println(s)
}

func init() {
	installCommand.StringVar(&user, "user", "", "user account to run the service")
	installCommand.StringVar(&workingdir, "workingdir", "", "working directory of the service")
	installCommand.StringVar(&config, "config", "/etc/myapp/config.ini", "config file of the service")
	runCommand.StringVar(&config, "config", defaultConfig, "config file of the service")
	flag.Usage = usage
}

func main() {
	var err error
	n := len(os.Args)
	if n <= 1 {
		fmt.Printf("invalid args\n")
		flag.Usage()
		return
	}

	subCmd := os.Args[1] // the second arg

	// get Config
	c, err := getServiceConfig(subCmd)
	if err != nil {
		fmt.Printf("get service config error: %s\n", err)
		return
	}

	prg := &NullService{}
	srv, err := service.New(prg, c)
	if err != nil {
		fmt.Printf("new service error: %s\n", err)
		return
	}

	err = runServiceControl(srv, subCmd)
	if err != nil {
		fmt.Printf("%s operation error: %s\n", subCmd, err)
		return
	}

	fmt.Printf("%s operation ok\n", subCmd)
	return
}

func runServiceControl(srv service.Service, subCmd string) error {
	switch subCmd {
	case "run":
		return run(config)
	default:
		return service.Control(srv, subCmd)
	}
}

func run(config string) error {
	// load info from config
	c, err := loadConfigFromIni(config)
	if err != nil {
		return err
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s]: receive a request from: %s\n", c.Server.Addr, r.RemoteAddr)
		w.Write([]byte("Welcome"))
	})
	fmt.Printf("listen on %s\n", c.Server.Addr)
	return http.ListenAndServe(c.Server.Addr, nil)
}

type NullService struct{}

func (p *NullService) Start(s service.Service) error {
	return nil
}

func (p *NullService) Stop(s service.Service) error {
	return nil
}

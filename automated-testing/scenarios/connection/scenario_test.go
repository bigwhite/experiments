package connection

import (
	"flag"
	"log"
	"os"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "", "the broker address(ip:port)")
}

func TestMain(m *testing.M) {
	flag.Parse()

	// setup for this scenario
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	/*
	   mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	   mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	   mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)
	*/

	// run this scenario test
	r := m.Run()

	// teardown for this scenario
	// tbd if teardown is needed

	os.Exit(r)
}

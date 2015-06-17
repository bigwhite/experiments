package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"sourcegraph.com/sourcegraph/appdash"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
)

var c Server

func init() {
	c = Server{
		CollectorAddr: ":3001",
		HTTPAddr:      ":3000",
	}
}

// Server is the command for running Appdash in server mode, where a
// collector server and the web UI are hosted.
type Server struct {
	CollectorAddr string
	HTTPAddr      string
}

func main() {
	var (
		memStore = appdash.NewMemoryStore()
		Store    = appdash.Store(memStore)
		Queryer  = memStore
	)

	app := traceapp.New(nil)
	app.Store = Store
	app.Queryer = Queryer

	var h http.Handler = app
	var l net.Listener
	var proto string
	var err error
	l, err = net.Listen("tcp", c.CollectorAddr)
	if err != nil {
		log.Fatal(err)
	}
	proto = "plaintext TCP (no security)"
	log.Printf("appdash collector listening on %s (%s)",
		c.CollectorAddr, proto)
	cs := appdash.NewServer(l, appdash.NewLocalCollector(Store))
	go cs.Start()

	log.Printf("appdash HTTP server listening on %s", c.HTTPAddr)
	err = http.ListenAndServe(c.HTTPAddr, h)
	if err != nil {
		fmt.Println("listenandserver listen err:", err)
	}
}

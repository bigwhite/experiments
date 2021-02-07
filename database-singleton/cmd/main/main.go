package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bigwhite/testdboper/pkg/config"
	"github.com/bigwhite/testdboper/pkg/db"
	"github.com/bigwhite/testdboper/pkg/reader"
	"github.com/bigwhite/testdboper/pkg/updater"
)

func init() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var quit = make(chan struct{})

	// do some init from db
	_ = db.DB()

	go func() {
		updater.Run(quit)
		wg.Done()
	}()
	go func() {
		reader.Run(quit)
		wg.Done()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	_ = <-c
	close(quit)
	log.Printf("recv exit signal...")
	wg.Wait()
	log.Printf("program exit ok")
}

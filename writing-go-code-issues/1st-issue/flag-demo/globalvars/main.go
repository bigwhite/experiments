package main

import (
	"flag"
	"fmt"
	"time"

	"./config"
	"./etcd"
)

func init() {
	flag.StringVar(&config.Endpoints, "endpoints", "127.0.0.1:2379", "comma-separated list of etcdv3 endpoints")
	flag.StringVar(&config.User, "user", "", "etcdv3 client user")
	flag.StringVar(&config.Password, "password", "", "etcdv3 client password")
}

func usage() {
	fmt.Println("flagdemo-app is a daemon application which provides xxx service.\n")
	fmt.Println("Usage of flagdemo-app:\n")
	fmt.Println("\t flagdemo-app [options]\n")
	fmt.Println("The options are:\n")

	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	go etcd.EtcdProxy()

	time.Sleep(5 * time.Second)
}

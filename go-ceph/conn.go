package main

import (
	"fmt"

	"github.com/ceph/go-ceph/rados"
)

func main() {
	conn, err := rados.NewConn()
	if err != nil {
		fmt.Println("error when invoke a new connection:", err)
		return
	}

	err = conn.ReadDefaultConfigFile()
	if err != nil {
		fmt.Println("error when read default config file:", err)
		return
	}

	err = conn.Connect()
	if err != nil {
		fmt.Println("error when connect:", err)
		return
	}

	fmt.Println("connect ceph cluster ok!")
	conn.Shutdown()
}

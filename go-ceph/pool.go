package main

import (
	"fmt"
	"os"

	"github.com/ceph/go-ceph/rados"
)

func newConn() (*rados.Conn, error) {
	conn, err := rados.NewConn()
	if err != nil {
		return nil, err
	}

	err = conn.ReadDefaultConfigFile()
	if err != nil {
		return nil, err
	}

	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func listPools(conn *rados.Conn, prefix string) {
	pools, err := conn.ListPools()
	if err != nil {
		fmt.Println("error when list pool", err)
		os.Exit(1)
	}
	fmt.Println(prefix, ":", pools)
}

func main() {
	conn, err := newConn()
	if err != nil {
		fmt.Println("error when invoke a new connection:", err)
		return
	}
	defer conn.Shutdown()
	fmt.Println("connect ceph cluster ok!")

	listPools(conn, "before make new pool")

	err = conn.MakePool("new_pool")
	if err != nil {
		fmt.Println("error when make new_pool", err)
		return
	}
	listPools(conn, "after make new pool")

	err = conn.DeletePool("new_pool")
	if err != nil {
		fmt.Println("error when delete pool", err)
		return
	}

	listPools(conn, "after delete new_pool")
}

package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

func socketFD(conn net.Conn) int {
	tcpConn := reflect.ValueOf(conn).Elem().FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := fdVal.Elem().FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", err)
			return
		}

		fmt.Printf("conn fd is [%d]\n", socketFD(conn))
	}
}

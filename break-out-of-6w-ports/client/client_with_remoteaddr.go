package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var remoteIP string

func init() {
	flag.StringVar(&remoteIP, "rip", "", "remoteIP")
}

func main() {
	flag.Parse()
	var count = 25000
	for i := 0; i < count; i++ {
		go func() {
			conn, err := net.Dial("tcp", remoteIP+":9000")
			if err != nil {
				fmt.Println("net.Dial error:", err)
				return
			}

			for {
				_, err := conn.Write([]byte("ping"))
				if err != nil {
					fmt.Println("conn.Write error:", err)
					return
				}
				time.Sleep(100 * time.Second)
			}
		}()
	}
	select {}
}

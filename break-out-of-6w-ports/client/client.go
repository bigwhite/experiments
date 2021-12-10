package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var count = 25000
	for i := 0; i < count; i++ {
		go func() {
			conn, err := net.Dial("tcp", "192.168.49.6:9000")
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

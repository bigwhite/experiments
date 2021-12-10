package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("listen ok")
	var mu sync.Mutex
	var count int

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error accept:", err)
			return
		}

		fmt.Printf("recv conn from [%s]\n", conn.RemoteAddr())
		go func(conn net.Conn) {
			var b = make([]byte, 10)
			for {
				_, err := conn.Read(b)
				if err != nil {
					e, ok := err.(net.Error)
					if ok {
						if e.Timeout() {
							continue
						}
					}

					mu.Lock()
					count--
					mu.Unlock()
					return
				}
				//println("recv: ", string(b))
			}
		}(conn)

		mu.Lock()
		count++
		mu.Unlock()
		fmt.Println("total count =", count)
	}

	select {}
}

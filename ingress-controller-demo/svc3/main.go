package main

import (
	"log"
	"net"
)

func echo(conn net.Conn) {
	for {
		buf := make([]byte, 256)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println("read error:", err)
			conn.Close()
			return
		}
		log.Printf("conn:[%v] recv data: %s", conn, string(buf))
		conn.Write(buf)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Println("error listening:", err.Error())
		return
	}
	defer l.Close()
	log.Println("listen ok")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("error accept:", err)
			return
		}

		log.Println("accept conn ok: ", conn)
		go echo(conn)
	}
}

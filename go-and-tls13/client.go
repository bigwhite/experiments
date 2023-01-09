package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:8443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello, world from client\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}

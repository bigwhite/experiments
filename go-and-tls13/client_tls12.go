package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS12,
	}

	conn, err := tls.Dial("tcp", "localhost:8443", conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		fmt.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(n, err)
		return
	}

	println(string(buf[:n]))
}

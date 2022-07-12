package main

import (
	"fmt"
	"time"

	tls "github.com/tjfoc/gmsm/gmtls"
)

func main() {
	cert, err := tls.LoadX509KeyPair("./certs/gm-client-enc-cert.pem", "./certs/gm-client-enc-key.pem")
	if err != nil {
		fmt.Println("load x509 keypair error:", err)
		return
	}

	conn, err := tls.Dial("tcp", "example.com:18000", &tls.Config{
		Certificates: []tls.Certificate{cert},
	})
	if err != nil {
		fmt.Println("failed to connect: " + err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("connect ok")
	for i := 0; i < 100; i++ {
		_, err := conn.Write([]byte("hello, gm"))
		if err != nil {
			fmt.Println("conn write error:", err)
			return
		}

		var b = make([]byte, 16)
		_, err = conn.Read(b)
		if err != nil {
			fmt.Println("conn read error:", err)
			return
		}
		fmt.Println(string(b))
		time.Sleep(time.Second)
	}
}

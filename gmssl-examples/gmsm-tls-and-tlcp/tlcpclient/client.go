package main

import (
	"fmt"
	"io/ioutil"
	"time"

	tls "github.com/tjfoc/gmsm/gmtls"
	x509 "github.com/tjfoc/gmsm/x509"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "./certs/ca-gm-cert.pem"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("read ca err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cert, err := tls.LoadX509KeyPair("./certs/client-gm-auth-cert.pem", "./certs/client-gm-auth-key.pem")
	if err != nil {
		fmt.Println("load x509 keypair error:", err)
		return
	}

	conn, err := tls.Dial("tcp", "example.com:18000", &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            pool,
		GMSupport:          tls.NewGMSupport(),
		InsecureSkipVerify: false,
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
		fmt.Println("reply:", string(b))
		time.Sleep(time.Second)
	}
}

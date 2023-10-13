package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//caCert, err := ioutil.ReadFile("ca-cert.pem")
	caCert, err := ioutil.ReadFile("inter-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	config := &tls.Config{
		RootCAs: caCertPool,
	}

	conn, err := tls.Dial("tcp", "server.com:8443", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 每秒发送信息
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		msg := "hello, tls"
		conn.Write([]byte(msg))

		// 读取回复
		buf := make([]byte, len(msg))
		conn.Read(buf)
		log.Println(string(buf))
	}

}

package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	serverAddr := flag.String("server", "server.com:8443", "Server address")
	flag.Parse()

	caCert, err := ioutil.ReadFile("inter-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	config := &tls.Config{
		RootCAs: caCertPool,
	}

	conn, err := tls.Dial("tcp", *serverAddr, config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 解析连接的服务器证书
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		log.Println("Server CN:", certs[0].Subject.CommonName)
	}

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

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// 1. 读取客户端证书文件
	clientCert, err := tls.LoadX509KeyPair("client-cert.pem", "client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// 2. 读取中间CA证书文件
	caCert, err := os.ReadFile("inter-cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// 3. 发送请求

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{clientCert},
				RootCAs:      certPool,
			},
		},
	}

	req, err := http.NewRequest("GET", "https://server.com:8443", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 4. 打印响应信息
	fmt.Println("Response Status:", resp.Status)
	//	fmt.Println("Response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

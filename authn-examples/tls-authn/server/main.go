package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	var validClients = map[string]struct{}{
		"client.com": struct{}{},
	}

	// 1. 加载证书文件
	cert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCert, err := os.ReadFile("inter-cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// 2. 配置TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // will trigger the invoke of VerifyPeerCertificate
		ClientCAs:    certPool,
	}

	// tls.Config设置
	tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		// 获取客户端证书
		cert := verifiedChains[0][0]

		// 提取CN作为客户端标识
		clientID := cert.Subject.CommonName
		fmt.Println(clientID)

		_, ok := validClients[clientID]
		if !ok {
			return errors.New("invalid client id")
		}

		return nil
	}
	// 添加处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// 3. 创建服务器
	srv := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	// 4. 启动服务器
	err = srv.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}

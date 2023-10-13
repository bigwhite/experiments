package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

// 服务端
func startServer(certFiles, keyFiles []string) {
	// 读取证书和密钥

	var certs []tls.Certificate
	for i := 0; i < len(certFiles); i++ {
		cert, err := tls.LoadX509KeyPair(certFiles[i], keyFiles[i])
		if err != nil {
			log.Fatal(err)
		}
		certs = append(certs, cert)
	}

	// 创建TLS配置
	config := &tls.Config{
		Certificates: certs,
	}

	// 启动TLS服务器
	listener, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// 处理连接...
	// 循环读取客户端的数据
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			// 读取失败则退出
			return
		}

		// 回显数据给客户端
		s := string(buf[:n])
		fmt.Printf("recv data: %s\n", s)
		conn.Write(buf[:n])
	}
}

func main() {
	certFiles := []string{"leaf-server-cert.pem", "leaf-server1-cert.pem", "leaf-server2-cert.pem"}
	keyFiles := []string{"leaf-server-key.pem", "leaf-server1-key.pem", "leaf-server2-key.pem"}

	// 启动服务器
	startServer(certFiles, keyFiles)
}

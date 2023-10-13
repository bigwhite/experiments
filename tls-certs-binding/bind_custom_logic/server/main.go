package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

// 服务端
func startServer(certsPath string) {

	// 创建TLS配置
	config := &tls.Config{
		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			// 根据clientHello信息选择cert

			certFile := fmt.Sprintf("%s/leaf-%s-cert.pem", certsPath, info.ServerName[:len(info.ServerName)-4])
			keyFile := fmt.Sprintf("%s/leaf-%s-key.pem", certsPath, info.ServerName[:len(info.ServerName)-4])

			// 读取证书和密钥
			cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			return &cert, err
		},
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
	// 启动服务器
	startServer("./certs")
}

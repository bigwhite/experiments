package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听指定的端口
	listener, err := net.Listen("tcp", "localhost:9004")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 9004")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // 处理连接
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // 确保连接在结束时关闭
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer) // 读取数据
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Printf("Recv: [%s]\n", string(buffer[:n]))

		// 将接收到的数据回显给客户端
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

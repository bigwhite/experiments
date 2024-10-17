package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	for {
		// 连接到服务器
		conn, err := net.Dial("tcp", "localhost:5000")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			time.Sleep(1 * time.Second) // 连接失败后等待再试
			continue
		}

		message := "Hello, Server!"
		_, err = conn.Write([]byte(message)) // 发送消息
		if err != nil {
			fmt.Println("Error sending message:", err)
			conn.Close()
			continue
		}

		// 读取回显的消息
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
		} else {
			fmt.Println("Received from server:", string(buffer[:n]))
		}

		conn.Close() // 完成后关闭连接

		time.Sleep(1 * time.Second) // 每秒发送一次
	}
}

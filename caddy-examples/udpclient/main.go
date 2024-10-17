package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 设置服务器地址
	address := net.UDPAddr{
		Port: 5000,
		IP:   net.ParseIP("localhost"),
	}

	for {
		// 每次发送消息前创建新的 UDP 连接
		conn, err := net.DialUDP("udp", nil, &address)
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			time.Sleep(1 * time.Second) // 连接失败后等待再试
			continue
		}

		message := []byte("Hello, UDP Server!")
		_, err = conn.Write(message) // 发送消息
		if err != nil {
			fmt.Println("Error sending message:", err)
			conn.Close()
			continue
		}

		// 读取回显的消息
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
		} else {
			fmt.Println("Received from server:", string(buffer[:n]))
		}

		conn.Close() // 完成后关闭连接

		time.Sleep(1 * time.Second) // 每秒发送一次
	}
}

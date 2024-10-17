package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听指定的 UDP 端口
	address := net.UDPAddr{
		Port: 9005,
		IP:   net.ParseIP("localhost"),
	}

	conn, err := net.ListenUDP("udp", &address)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("UDP server is listening on port 9005")

	buffer := make([]byte, 1024)

	for {
		// 接收数据
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		fmt.Printf("Received from %s: %s\n", addr.String(), string(buffer[:n]))

		// 将接收到的数据回显给客户端
		_, err = conn.WriteToUDP(buffer[:n], addr)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
			continue
		}
	}
}

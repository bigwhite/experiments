package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: nat-hole-puncher <local_ip> <local_port> <target_ip> <target_port>")
		return
	}

	localIP := os.Args[1]
	localPort := os.Args[2]
	targetIP := os.Args[3]
	targetPort := os.Args[4]

	// 向target_ip:target_port发送数据
	err := sendUDPMessage("Hello, World!", localIP, localPort, targetIP+":"+targetPort)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Println("sending message to", targetIP+":"+targetPort, "ok")

	// 向target_ip:target_port+1发送数据
	p, _ := strconv.Atoi(targetPort)
	nextTargetPort := fmt.Sprintf("%d", p+1)
	err = sendUDPMessage("Hello, World!", localIP, localPort, targetIP+":"+nextTargetPort)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Println("sending message to", targetIP+":"+nextTargetPort, "ok")

	// 重新监听local addr
	startUDPReceiver(localIP, localPort)
}

func sendUDPMessage(message, localIP, localPort, target string) error {
	addr, err := net.ResolveUDPAddr("udp", target)
	if err != nil {
		return err
	}

	lport, _ := strconv.Atoi(localPort)
	conn, err := net.DialUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(localIP),
		Port: lport,
	}, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// 发送数据
	_, err = conn.Write([]byte(message))
	if err != nil {
		return err
	}

	return nil
}

func startUDPReceiver(ip, port string) {
	addr, err := net.ResolveUDPAddr("udp", ip+":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()
	fmt.Println("listen address:", ip+":"+port, "ok")

	buf := make([]byte, 1024)
	for {
		n, senderAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Received message: %s from %s\n", string(buf[:n]), senderAddr.String())
	}
}

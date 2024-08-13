package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: udp-client-addr-display <local_ip> <local_port>")
		return
	}

	localIP := os.Args[1]
	localPort := os.Args[2]

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		startUDPReceiver(localIP, localPort)
	}()

	go func() {
		defer wg.Done()
		p, _ := strconv.Atoi(localPort)
		nextLocalPort := fmt.Sprintf("%d", p+1)
		startUDPReceiver(localIP, nextLocalPort)
	}()

	wg.Wait()
}

func startUDPReceiver(localIP, localPort string) {
	addr, err := net.ResolveUDPAddr("udp", localIP+":"+localPort)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	n, clientAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Received message: %s from %s\n", string(buf[:n]), clientAddr.String())
}

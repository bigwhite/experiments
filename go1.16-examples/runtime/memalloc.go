package main

import "time"

func allocMem() []byte {
	b := make([]byte, 1024*1024*1) //1M
	return b
}

func main() {
	for i := 0; i < 100000; i++ {
		_ = allocMem()
		time.Sleep(500 * time.Millisecond)
	}
}

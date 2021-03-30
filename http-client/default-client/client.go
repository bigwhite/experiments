package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(256)
	for i := 0; i < 256; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}()
	}
	wg.Wait()
}

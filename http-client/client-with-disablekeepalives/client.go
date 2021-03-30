package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	tr := &http.Transport{
		MaxConnsPerHost:     5,
		MaxIdleConnsPerHost: 3,
		IdleConnTimeout:     10 * time.Second,
		DisableKeepAlives:   true,
	}
	client := http.Client{
		Transport: tr,
	}
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := client.Get("http://localhost:8080")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			fmt.Printf("g-%d: %s\n", i, string(body))
		}(i)
	}
	wg.Wait()

	time.Sleep(5 * time.Second)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for i := 0; i < 2; i++ {
				resp, err := client.Get("http://localhost:8080")
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				fmt.Printf("g-%d: %s\n", i+10, string(body))
				time.Sleep(time.Second)
			}
		}(i)
	}

	time.Sleep(15 * time.Second)
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				resp, err := client.Get("http://localhost:8080")
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				fmt.Printf("g-%d: %s\n", i+20, string(body))
				time.Sleep(time.Second)
			}
		}(i)
	}
	wg.Wait()
}

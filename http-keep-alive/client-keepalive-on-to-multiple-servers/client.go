package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{}
	req1, err := http.NewRequest("Get", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	req2, err := http.NewRequest("Get", "http://localhost:8081", nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		resp1, err := c.Do(req1)
		if err != nil {
			fmt.Println("http get error:", err)
			return
		}
		defer resp1.Body.Close()

		b1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Println("read body error:", err)
			return
		}
		log.Println("response1 body:", string(b1))

		resp2, err := c.Do(req2)
		if err != nil {
			fmt.Println("http get error:", err)
			return
		}
		defer resp2.Body.Close()

		b2, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			fmt.Println("read body error:", err)
			return
		}
		log.Println("response2 body:", string(b2))

		time.Sleep(5 * time.Second)
	}

}

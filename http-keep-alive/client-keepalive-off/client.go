package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	c := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("Get", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("http get error:", err)
			return
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read body error:", err)
			return
		}
		log.Println("response body:", string(b))
		time.Sleep(5 * time.Second)
	}

}

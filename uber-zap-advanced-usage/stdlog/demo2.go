package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./demo2.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.Println("this is go standard log package")
}

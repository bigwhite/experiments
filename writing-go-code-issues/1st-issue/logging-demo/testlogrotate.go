package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("./app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	defer file.Close()

	logger := log.New(file,
		"APP_LOG_PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	for {
		logger.Println("test log")
		time.Sleep(time.Second * 1)
	}
}

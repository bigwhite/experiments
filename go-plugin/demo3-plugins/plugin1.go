package main

import (
	"log"

	_ "github.com/go-redis/redis/v8"
	_ "github.com/sirupsen/logrus"
)

func init() {
	log.Println("plugin1 init")
}

package common

import (
	"fmt"
	"log"
)

func init() {
	log.Println("common init")
}

func Bar() {
	fmt.Println("bar in common")
}

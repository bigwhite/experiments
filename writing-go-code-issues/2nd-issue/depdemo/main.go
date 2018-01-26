package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Person struct {
	ID   string
	Name string
	City string
}

func main() {
	p := Person{}
	if _, err := toml.DecodeFile("./data.toml", &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}

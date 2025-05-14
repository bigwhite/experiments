package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
type Person struct {
	Name    string  `json:"name"`
	Address Address `json:",inline"` // v2 支持
}

func main() {
	fmt.Println("--- Testing 'inline' Tag ---")
	p := Person{
		Name:    "Tony Bai",
		Address: Address{Street: "123 Go Ave", City: "Gopher City"},
	}
	jsonData, _ := json.Marshal(p, json.Deterministic(true))
	(*jsontext.Value)(&jsonData).Indent() // indent for readability
	fmt.Println("Serialized Person (v2 expected with inline):\n", string(jsonData))
}

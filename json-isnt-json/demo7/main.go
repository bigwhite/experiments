package main

import (
	"encoding/json" // 在jsonv2时，改为"encoding/json/v2"
	"fmt"
)

func main() {
	var data map[string]int

	// Duplicate keys - last value wins (no error)
	err := json.Unmarshal([]byte(`{"a": 1, "a": 2}`), &data)
	if err != nil {
		fmt.Println("Duplicate key error:", err)
	} else {
		fmt.Printf("Duplicate keys allowed, value: %d\n", data["a"]) // 2
	}

	// Trailing commas - error
	err = json.Unmarshal([]byte(`{"a": 1,}`), &data)
	if err != nil {
		fmt.Println("Trailing comma error:", err)
	}

	// Leading zeros - error
	err = json.Unmarshal([]byte(`{"num": 007}`), &data)
	if err != nil {
		fmt.Println("Leading zeros error:", err)
	}

	// Single quotes - error
	err = json.Unmarshal([]byte(`{'a': 1}`), &data)
	if err != nil {
		fmt.Println("Single quotes error:", err)
	}
}

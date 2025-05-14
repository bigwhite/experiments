package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Enabled     bool    `json:"enabled,omitempty"`     // v1: false 时省略; v2: false 不编码为JSON空则不省略
	Count       int     `json:"count,omitempty"`       // v1: 0 时省略; v2: 0 不编码为JSON空则不省略
	Name        string  `json:"name,omitempty"`        // v1 & v2: "" 时省略
	Description *string `json:"description,omitempty"` // v1 & v2: nil 时省略

	IsSet  bool    `json:"is_set,omitzero"`  // v1(1.24+)/v2: false 时省略
	Port   int     `json:"port,omitzero"`    // v1(1.24+)/v2: 0 时省略
	APIKey *string `json:"api_key,omitzero"` // v1(1.24+)/v2: nil 时省略
}

func main() {
	fmt.Println("--- Testing omitempty/omitzero ---")
	emptyConf := Config{} // All zero values
	descValue := ""
	emptyConfWithEmptyStringPtr := Config{Description: &descValue, APIKey: &descValue}

	jsonDataV1, _ := json.MarshalIndent(emptyConf, "", "  ")
	fmt.Println("V1 (go run) - Empty Config:\n", string(jsonDataV1))
	jsonDataV1Ptr, _ := json.MarshalIndent(emptyConfWithEmptyStringPtr, "", "  ")
	fmt.Println("V1 (go run) - Empty Config with Empty String Ptr:\n", string(jsonDataV1Ptr))
}

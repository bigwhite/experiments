package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
)

type Item struct {
	ID            string                    `json:"id"`
	KnownData     string                    `json:"known_data"`
	UnknownFields map[string]jsontext.Value `json:",unknown"`
}

func main() {
	fmt.Println("--- Testing 'unknown' Tag ---")
	inputJSON := `{"id":"item1","known_data":"some data","new_field":"value for new field","another_unknown":123, "obj_field":{"nested":true}}`
	var item Item
	err := json.Unmarshal([]byte(inputJSON), &item)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Unmarshaled Item: %+v\n", item)
	if item.UnknownFields != nil {
		fmt.Println("Captured Unknown Fields:")
		for k, v := range item.UnknownFields {
			fmt.Printf("  %s: %s\n", k, string(v))
		}
	}
}

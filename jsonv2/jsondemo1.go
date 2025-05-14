package main

import (
	"encoding/json"
	"fmt"
)

type TargetRepeat struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("--- Testing Duplicate Keys ---")
	inputJSONRepeat := `{"message": "hello 1", "message": "hello 2"}` // 重复键 "message"

	var outRepeat TargetRepeat
	errRepeat := json.Unmarshal([]byte(inputJSONRepeat), &outRepeat)
	if errRepeat != nil {
		fmt.Println("Unmarshal with duplicate keys error (expected for v2):", errRepeat)
	} else {
		fmt.Printf("Unmarshal with duplicate keys output (v1 behavior): %+v\n", outRepeat)
	}

	fmt.Println("\n--- Testing Case Sensitivity ---")
	type TargetCase struct {
		MyValue string `json:"myValue"` // Tag is camelCase
	}
	inputJSONCase := `{"myvalue": "hello case"}` // JSON key is lowercase

	var outCase TargetCase
	errCase := json.Unmarshal([]byte(inputJSONCase), &outCase)
	if errCase != nil {
		fmt.Println("Unmarshal with case mismatch error (expected for v2 default):", errCase)
	} else {
		fmt.Printf("Unmarshal with case mismatch output (v1 behavior or v2 with nocase): %+v\n", outCase)
		if outCase.MyValue == "" {
			fmt.Println("Note: myValue field was not populated due to case mismatch in v2 (default).")
		}
	}
}

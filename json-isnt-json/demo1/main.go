package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"id": 9007199254740993}`)
	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)
	// id 被解析为 float64，精度丢失！
	fmt.Printf("v1 with interface{}: %.0f\n", data["id"]) // 输出: 9007199254740992

	var typed struct {
		ID int64 `json:"id"`
	}
	json.Unmarshal(jsonData, &typed)
	fmt.Println(typed.ID) // 输出: 9007199254740993
}


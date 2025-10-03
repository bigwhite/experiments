package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 场景一：description 显式为 null
	jsonWithNull := []byte(`{"name":"Gopher", "description":null}`)
	// 场景二：description 字段缺失
	jsonMissing := []byte(`{"name":"Gopher"}`)

	distinguish(jsonWithNull)
	distinguish(jsonMissing)
}

func distinguish(jsonData []byte) {
	// 步骤 1: 解码到一个 map[string]json.RawMessage
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(jsonData, &raw); err != nil {
		panic(err)
	}

	// 步骤 2: 检查 "description" 键是否存在
	descData, ok := raw["description"]
	if !ok {
		fmt.Println("Result: 'description' key is MISSING.")
		return
	}

	// 步骤 3: 如果键存在，检查其内容是否为 "null"
	if string(descData) == "null" {
		fmt.Println("Result: 'description' key is explicitly NULL.")
		return
	}

	// 如果存在且不为 null，则可以进一步解码
	var desc string
	json.Unmarshal(descData, &desc)
	fmt.Printf("Result: 'description' has value: %s\n", desc)
}

package main

import (
	"encoding/json"
	"fmt"
)

type UserUpdatePayload struct {
	Nickname    string  `json:"nickname"`
	Description *string `json:"description"` // 指针字段表示可选
}

func main() {
	// 场景一：用户想将 description 更新为空字符串 ""
	jsonWithValue := []byte(`{"nickname":"Gopher", "description":""}`)
	var u1 UserUpdatePayload
	json.Unmarshal(jsonWithValue, &u1)
	fmt.Printf("Scenario 1 (Zero Value): Description is nil: %t, Value: '%s'\n", u1.Description == nil, *u1.Description)

	// 场景二：用户未提供 description 字段 (无论是显式 null 还是 missing)
	jsonWithoutValue := []byte(`{"nickname":"Gopher"}`) // or {"description":null}
	var u2 UserUpdatePayload
	json.Unmarshal(jsonWithoutValue, &u2)
	fmt.Printf("Scenario 2 (Absence):      Description is nil: %t\n", u2.Description == nil)
}

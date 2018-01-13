package main

import (
	"encoding/json"
	"fmt"
)

func marshalResponse(code int, msg string, result interface{}) (string, error) {
	m := map[string]interface{}{
		"code":   0,
		"msg":    "ok",
		"result": result,
	}

	b, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func marshalResponse1(code int, msg string, result interface{}) (string, error) {
	s, ok := result.(string)
	var m = map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}

	if ok {
		rawData := json.RawMessage(s)
		m["result"] = rawData
	} else {
		m["result"] = result
	}

	b, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	s, err := marshalResponse(0, "ok", `{"name": "tony", "city": "shenyang"}`)
	if err != nil {
		fmt.Println("marshal response error:", err)
		return
	}
	fmt.Println(s)

	s, err = marshalResponse1(0, "ok", `{"name": "tony", "city": "shenyang"}`)
	if err != nil {
		fmt.Println("marshal response1 error:", err)
		return
	}
	fmt.Println(s)
}

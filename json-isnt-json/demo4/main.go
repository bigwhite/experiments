package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
	// --- 问题场景 ---
	// 假设两个不同的系统需要对相同的业务数据进行签名。

	// 系统 1: 一个 JavaScript 服务，它保留了对象属性的插入顺序。
	// 注意键的顺序: "currency" 在前, "amount" 在后。
	jsONString := `{"currency":"USD","amount":100}`
	fmt.Printf("JSON from JS-like system: %s\n", jsONString)

	// 系统 2: 一个 Go 服务，它序列化一个 map。
	data := map[string]interface{}{
		"currency": "USD",
		"amount":   100,
	}

	// Go 的 json.Marshal 会对 map 的键按字母顺序排序。
	// 因此 "amount" 会排在 "currency" 前面。
	goJSONBytes, _ := json.Marshal(data)
	goJSONString := string(goJSONBytes)
	fmt.Printf("JSON from Go system (map): %s\n", goJSONString)

	// --- 导致的后果: 加密签名失败 ---
	secret := []byte("my-super-secret-key")

	// 为 JS 风格的 JSON 字符串计算 HMAC
	hmacJS := calculateHMAC(secret, []byte(jsONString))
	fmt.Printf("HMAC for JS JSON: %s\n", hmacJS)

	// 为 Go 生成的 JSON 字符串计算 HMAC
	hmacGo := calculateHMAC(secret, goJSONBytes)
	fmt.Printf("HMAC for Go JSON: %s\n", hmacGo)

	// 比较两个签名
	signaturesMatch := hmac.Equal([]byte(hmacJS), []byte(hmacGo))
	fmt.Printf("\nDo the signatures match? %t\n", signaturesMatch)
	if !signaturesMatch {
		fmt.Println("Authentication Fails! The byte representations were different.")
	}
}

// calculateHMAC 是一个辅助函数，用于计算并编码 HMAC 值
func calculateHMAC(secret, data []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

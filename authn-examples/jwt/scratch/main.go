package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Claims struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Iat  int64  `json:"iat"`
}

// GenerateToken：不依赖第三方库的JWT生成实现
func GenerateToken(claims *Claims, key string) (string, error) {
	header, _ := json.Marshal(Header{
		Alg: "HS256",
		Typ: "JWT",
	})
	// 序列化Payload
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// 拼接成JWT字符串
	headerEncoded := base64.RawURLEncoding.EncodeToString(header)
	payloadEncoded := base64.RawURLEncoding.EncodeToString([]byte(payload))

	encodedToSign := headerEncoded + "." + payloadEncoded

	// 使用HMAC+SHA256签名
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(encodedToSign))
	sig := hash.Sum(nil)
	sigEncoded := base64.RawURLEncoding.EncodeToString(sig)

	var token string
	token += headerEncoded
	token += "."
	token += payloadEncoded
	token += "."
	token += sigEncoded

	return token, nil
}

func main() {
	var claims = &Claims{
		Sub:  "1234567890",
		Name: "John Doe",
		Iat:  1516239022,
	}

	result, _ := GenerateToken(claims, "iamtonybai")
	fmt.Println(result)
}

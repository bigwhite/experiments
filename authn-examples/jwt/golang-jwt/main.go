package main

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Sub                  string `json:"sub"`
	Name                 string `json:"name"`
	jwt.RegisteredClaims        // use its Subject and IssuedAt
}

func main() {
	mySigningKey := []byte("iamtonybai")

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		Name: "John Doe",
		Sub:  "1234567890",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Unix(1516239022, 0)), //  1516239022
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Println(ss)

	_, err := verifyToken(ss, "iamtonybai")
	if err != nil {
		fmt.Println("invalid token:", err)
		return
	}

	fmt.Println("valid token")
}

// verifyToken 验证JWT函数
func verifyToken(tokenString, key string) (*jwt.Token, error) {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证签名
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

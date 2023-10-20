package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// 创建一个基本的HTTP服务器
	mux := http.NewServeMux()

	username := "admin"
	password := "123456"
	key := "iamtonybai"

	// 针对/的handler
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// 返回401 Unauthorized响应
		w.Header().Set("WWW-Authenticate", "Basic realm=\"server.com\"")
		w.WriteHeader(http.StatusUnauthorized)
	})

	// login handler
	mux.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		// 从请求头中获取Basic Auth认证信息
		user, pass, ok := req.BasicAuth()
		if !ok {
			// 认证失败
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// 验证用户名密码
		if user == username && pass == password {
			// 认证成功，生成token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": username,
				"iat":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			})
			signedToken, _ := token.SignedString([]byte(key))
			w.Write([]byte(signedToken))
		} else {
			// 认证失败
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	})

	// calc handler
	mux.HandleFunc("/calc", func(w http.ResponseWriter, req *http.Request) {
		// 读取并校验jwt token
		token := req.Header.Get("Authorization")[len("Bearer "):]
		fmt.Println(token)
		if _, err := verifyToken(token, key); err != nil {
			// 认证失败
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		w.Write([]byte("invoke calc ok"))
	})

	// 监听8080端口
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
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

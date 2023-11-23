package main

import (
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

	// for uptime test
	mux.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// login handler
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
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
				"iat":      jwt.NewNumericDate(time.Now()),
			})
			signedToken, _ := token.SignedString([]byte(key))
			w.Write([]byte(signedToken))
		} else {
			// 认证失败
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	})

	// 监听28084端口
	err := http.ListenAndServe(":28084", mux)
	if err != nil {
		log.Fatal(err)
	}
}

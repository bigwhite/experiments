package main

import (
	"log"
	"net/http"
)

func main() {
	// 创建一个基本的HTTP服务器
	mux := http.NewServeMux()

	username := "admin"
	password := "123456"

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
			// 认证成功
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Welcome to the protected resource!"))
		} else {
			// 认证失败
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	})

	// 监听8080端口
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

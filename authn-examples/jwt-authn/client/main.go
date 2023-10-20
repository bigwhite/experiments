package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://server.com:8080/", nil)

	// 发送默认请求
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析响应头
	authHeader := response.Header.Get("WWW-Authenticate")
	loginReq, _ := http.NewRequest("POST", "http://server.com:8080/login", nil)
	username := "admin"
	password := "123456"

	// 判断认证类型
	if !strings.Contains(authHeader, "Basic") {
		// 不支持的认证类型
		fmt.Println("Unsupported authentication type:", authHeader)
		return
	}

	// 使用Basic Auth, 添加Basic Auth头
	loginReq.SetBasicAuth(username, password)
	response, err = client.Do(loginReq)

	fmt.Println(response.StatusCode)

	// 从响应包体中获取服务端分配的jwt token
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := string(body)
	fmt.Println("token=", token)

	// 基于token访问服务端其他功能
	apiReq, _ := http.NewRequest("POST", "http://server.com:8080/calc", nil)
	apiReq.Header.Set("Authorization", "Bearer "+token)
	response, err = client.Do(apiReq)
	fmt.Println(response.StatusCode)
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

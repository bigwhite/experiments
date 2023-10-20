package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
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
	if !strings.Contains(authHeader, "Digest") {
		// 不支持的认证类型
		fmt.Println("Unsupported authentication type:", authHeader)
		return
	}

	// 使用Digest Auth

	//随机数
	cnonce := GenNonce()

	//生成HA1
	ha1 := GetHA1(username, password, cnonce)

	//构建Authorization头
	auth := "Digest username=\"" + username + "\", nonce=\"" + cnonce + "\", algorithm=MD5, response=\"" + GetResponse(ha1, cnonce) + "\""

	loginReq.Header.Set("Authorization", auth)
	response, err = client.Do(loginReq)

	// 打印响应状态
	fmt.Println(response.StatusCode)

	// 打印响应包体
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 生成随机数
func GenNonce() string {
	h := md5.New()
	io.WriteString(h, fmt.Sprint(rand.Int()))
	return hex.EncodeToString(h.Sum(nil))
}

// 根据用户名密码和随机数生成HA1
func GetHA1(username, password, cnonce string) string {
	h := md5.New()
	io.WriteString(h, username+":"+cnonce+":"+password)
	return hex.EncodeToString(h.Sum(nil))
}

// 根据HA1,随机数生成response
func GetResponse(ha1, cnonce string) string {
	h := md5.New()
	io.WriteString(h, strings.ToUpper("md5")+":"+ha1+":"+cnonce+"::"+strings.ToUpper("md5"))
	return hex.EncodeToString(h.Sum(nil))
}

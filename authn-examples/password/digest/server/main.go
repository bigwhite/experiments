package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	password := "123456"

	// 针对/的handler
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// 返回401 Unauthorized响应
		w.Header().Set("WWW-Authenticate", "Digest realm=\"server.com\"")
		w.WriteHeader(http.StatusUnauthorized)
	})

	// login handler
	mux.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Header)

		//验证参数
		if Verify(req, password) {
			fmt.Fprintln(w, "Verify Success!")
		} else {
			w.WriteHeader(401)
			fmt.Fprintln(w, "Verify Failed!")
		}
	})

	// 监听8080端口
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func Verify(r *http.Request, password string) bool {
	auth := r.Header.Get("Authorization")
	params := strings.Split(auth, ",")
	var username, cnonce, response string

	for _, p := range params {
		p := strings.Trim(p, " ")
		kv := strings.Split(p, "=")
		if kv[0] == "Digest username" {
			username = strings.Trim(kv[1], "\"")
		}
		if kv[0] == "nonce" {
			cnonce = strings.Trim(kv[1], "\"")
		}
		if kv[0] == "response" {
			response = strings.Trim(kv[1], "\"")
		}
	}

	if username == "" {
		return false
	}

	//根据用户名密码及随机数生成HA1
	ha1 := GetHA1(username, password, cnonce)

	//自己生成response与请求中response对比
	return response == GetResponse(ha1, cnonce)
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

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"text/template"
	"time"
)

func randString(n int) string {
	// 返回长度为n的随机字符串
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func portalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("portalHandler:", *r)

	// 获取请求参数用于渲染应答html页面
	clientID := r.FormValue("client_id")
	scopeTxt := r.FormValue("scope")
	state := r.FormValue("state")
	redirectURI := r.FormValue("redirect_uri")

	// 渲染授权页面模板
	tmpl := template.Must(template.ParseFiles("portal.html"))
	data := map[string]interface{}{
		"AppName":     clientID,
		"Scopes":      strings.Split(scopeTxt, ","),
		"ScopeTxt":    scopeTxt,
		"State":       state,
		"RedirectURI": redirectURI,
	}
	tmpl.Execute(w, data)
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authorizeHandler:", *r)

	responsTyp := r.FormValue("response_type")
	if responsTyp != "code" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := r.FormValue("username")
	password := r.FormValue("password")

	mu.Lock()
	v, ok := validUsers[user]
	if !ok {
		fmt.Println("not found the user:", user)
		mu.Unlock()
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		return
	}
	mu.Unlock()

	if v != password {
		fmt.Println("invalid password")
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		return
	}

	clientID := r.FormValue("client_id")
	scopeTxt := r.FormValue("scope")
	state := r.FormValue("state")
	redirectURI := r.FormValue("redirect_uri")

	code := randString(8)
	mu.Lock()
	codeCache[code] = authorizeContext{
		clientID:    clientID,
		scopeTxt:    scopeTxt,
		state:       state,
		redirectURI: redirectURI,
	}
	mu.Unlock()

	unescapeURI, _ := url.QueryUnescape(redirectURI)
	redirectURI = fmt.Sprintf("%s?code=%s&state=%s", unescapeURI, code, state)
	w.Header().Add("Location", redirectURI)
	w.WriteHeader(http.StatusFound)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tokenHandler:", *r)

	// check client_id and client_secret
	user, password, ok := r.BasicAuth()
	if !ok {
		fmt.Println("no authorization header")
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		return
	}

	mu.Lock()
	v, ok := validClients[user]
	if !ok {
		fmt.Println("not found user:", user)
		mu.Unlock()
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		return
	}
	mu.Unlock()

	if v != password {
		fmt.Println("invalid password")
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		return
	}

	// check code and redirect_uri
	code := r.FormValue("code")
	redirect_uri := r.FormValue("redirect_uri")

	mu.Lock()
	ac, ok := codeCache[code]
	if !ok {
		fmt.Println("not found code:", code)
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		return
	}
	mu.Unlock()

	if ac.redirectURI != redirect_uri {
		fmt.Println("invalid redirect_uri:", redirect_uri)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// generate access_token
	accessToken := randString(16)

	respTxt := fmt.Sprintf(`{ 
       "access_token":"%s",
       "expires_in":3600
	}`, accessToken)

	w.Write([]byte(respTxt))
}

var validUsers = map[string]string{
	"tonybai": "123456",
}

// client_id:client_secret
var validClients = map[string]string{
	"my-photo-print": "123456",
}

var mu sync.Mutex

type authorizeContext struct {
	clientID    string
	scopeTxt    string
	state       string
	redirectURI string
}

var codeCache = make(map[string]authorizeContext)

func main() {
	http.HandleFunc("/oauth/portal", portalHandler)
	http.HandleFunc("/oauth/authorize", authorizeHandler)
	http.HandleFunc("/oauth/token", tokenHandler)

	fmt.Println("启动授权服务器，监听8081端口")
	http.ListenAndServe(":8081", nil)
}

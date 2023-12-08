package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var mu sync.Mutex
var stateCache = map[string]struct{}{}
var userProfile = make(map[string]*Profile)

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

// 照片冲印主页，引导用户去授权平台认证
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler:", *r)

	// 渲染首页页面模板
	var state = randString(6)
	mu.Lock()
	stateCache[state] = struct{}{}
	mu.Unlock()
	tmpl := template.Must(template.ParseFiles("home.html"))
	data := map[string]interface{}{
		"State": state,
	}
	tmpl.Execute(w, data)
}

// callback handler，用户(EU)拿到code后调用该handler
func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oauthCallbackHandler:", *r)

	code := r.FormValue("code")
	state := r.FormValue("state")

	// check state
	mu.Lock()
	_, ok := stateCache[state]
	if !ok {
		mu.Unlock()
		fmt.Println("not found state:", state)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	delete(stateCache, state)
	mu.Unlock()

	// fetch access_token and id_token with code
	accessToken, idToken, err := fetchAccessTokenAndIDToken(code)
	if err != nil {
		fmt.Println("fetch access_token error:", err)
		return
	}
	fmt.Println("fetch access_token ok:", accessToken)

	// parse id_token
	mySigningKey := []byte("iamtonybai")
	claims := jwt.RegisteredClaims{}
	_, err = jwt.ParseWithClaims(idToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("parse id_token error:", err)
		return
	}

	// use access_token and userID to get user info
	up, err := getUserInfo(accessToken, claims.Subject)
	if err != nil {
		fmt.Println("get user info error:", err)
		return
	}
	fmt.Println("get user info ok:", up)

	mu.Lock()
	userProfile[claims.Subject] = up
	mu.Unlock()

	// 设置cookie
	cookie := http.Cookie{
		Name:   "my-photo-print.com-session",
		Value:  claims.Subject,
		Domain: "my-photo-print.com",
		Path:   "/profile",
	}
	http.SetCookie(w, &cookie)
	w.Header().Add("Location", "/profile")
	w.WriteHeader(http.StatusFound) // redirect to /profile
}

type Profile struct {
	Name     string
	Homepage string
	Mailbox  string
}

const authServer = "http://open.my-yunpan.com:8081"

var client_id = "my-photo-print"
var client_secert = "123456"

func fetchAccessTokenAndIDToken(code string) (string, string, error) {
	// 构建请求参数
	params := url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("code", code)
	params.Set("redirect_uri", "http://my-photo-print.com:8080/oauth/cb")

	serverURL := authServer + "/oauth/token"

	// 拼接带参数的URL
	fullURL := fmt.Sprintf("%s?%s", serverURL, params.Encode())

	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		return "", "", err
	}

	// 添加HTTP Basic Auth头信息
	req.SetBasicAuth(client_id, client_secert)

	// 发起POST请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("response status is %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	// 解析JSON响应
	var response struct {
		AccessToken string `json:"access_token"`
		IDToken     string `json:"id_token"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", "", err
	}

	return response.AccessToken, response.IDToken, nil
}

func getUserInfo(accessToken, userID string) (*Profile, error) {
	// 构建带有accessToken的URL
	serverURL := authServer + "/userinfo"
	url := fmt.Sprintf("%s?accessToken=%s&id=%s", serverURL, accessToken, userID)

	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p Profile
	_ = json.Unmarshal(body, &p)
	return &p, nil
}

// user profile页面
func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("profileHandler:", *r)

	cookie, err := r.Cookie("my-photo-print.com-session")
	if err != nil {
		http.Error(w, "找不到cookie，请重新登录", 401)
		return
	}
	fmt.Printf("found cookie: %#v\n", cookie)

	mu.Lock()
	pf, ok := userProfile[cookie.Value]
	if !ok {
		mu.Unlock()
		fmt.Println("not found user:", cookie.Value)
		// 跳转到首页
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	mu.Unlock()

	// 渲染照片页面模板
	tmpl := template.Must(template.ParseFiles("profile.html"))
	tmpl.Execute(w, pf)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/oauth/cb", oauthCallbackHandler)
	fmt.Println("照片冲印服务启动，监听8080端口")
	http.ListenAndServe(":8080", nil)
}

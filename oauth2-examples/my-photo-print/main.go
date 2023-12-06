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
)

var mu sync.Mutex
var stateCache = map[string]struct{}{}
var userPhotoList = make(map[string][]string)

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

// 照片冲印主页，引导用户去授权平台
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

// callback handler，用户拿到code后调用该handler
func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oauthCallbackHandler:", *r)

	code := r.FormValue("code")
	state := r.FormValue("state")

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

	// fetch access_token with code
	accessToken, err := fetchAccessToken(code)
	if err != nil {
		fmt.Println("fetch access_token error:", err)
		return
	}
	fmt.Println("fetch access_token ok:", accessToken)

	// use access_token to get user's photo list
	user, pl, err := getPhotoList(accessToken)
	if err != nil {
		fmt.Println("get photo list error:", err)
		return
	}
	fmt.Println("get photo list ok:", pl)

	mu.Lock()
	userPhotoList[user] = pl
	mu.Unlock()

	w.Header().Add("Location", "/photos?user="+user)
	w.WriteHeader(http.StatusFound)
}

const yunpanServer = "http://my-yunpan.com:8082"
const authServer = "http://open.my-yunpan.com:8081"

var client_id = "my-photo-print"
var client_secert = "123456"

func fetchAccessToken(code string) (string, error) {
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
		return "", err
	}

	// 添加HTTP Basic Auth头信息
	req.SetBasicAuth(client_id, client_secert)

	// 发起POST请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response status is %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析JSON响应
	var response struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.AccessToken, nil
}

func getPhotoList(accessToken string) (string, []string, error) {
	// 构建带有accessToken的URL
	serverURL := yunpanServer + "/photos"
	url := fmt.Sprintf("%s?accessToken=%s&method=listall", serverURL, accessToken)

	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	fmt.Println(string(body))
	m := make(map[string][]string)
	err = json.Unmarshal(body, &m)
	if err != nil {
		return "", nil, err
	}

	var user string
	var pl []string
	for u, v := range m {
		user = u
		pl = v
		break
	}
	return user, pl, nil
}

// 待获取到用户照片数据后，让用户浏览器重定向到该页面
func listPhonesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listPhonesHandler:", *r)

	user := r.FormValue("user")
	mu.Lock()
	pl, ok := userPhotoList[user]
	if !ok {
		mu.Unlock()
		fmt.Println("not found user:", user)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	mu.Unlock()

	// 渲染照片页面模板
	tmpl := template.Must(template.ParseFiles("photolist.html"))
	data := map[string]interface{}{
		"Username":  user,
		"PhotoList": pl,
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/photos", listPhonesHandler)
	http.HandleFunc("/oauth/cb", oauthCallbackHandler)
	fmt.Println("照片冲印服务启动，监听8080端口")
	http.ListenAndServe(":8080", nil)
}

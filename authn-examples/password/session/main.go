package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("session-key"))

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/calcAdd", calcAdd)

	http.ListenAndServe(":8080", nil)
}

var credentials = map[string]string{
	"admin": "123456",
	"test":  "654321",
}

func isValid(username, password string) bool {
	// 验证用户名密码
	v, ok := credentials[username]
	if !ok {
		return false
	}

	if v != password {
		return false
	}
	return true
}

func base64Encode(src string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(src))
	return encoded
}

func base64Decode(encoded string) string {
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	return string(decoded)
}

func randomStr() string {
	// 生成随机数
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100000)

	// 格式化为05位字符串
	str := fmt.Sprintf("%05d", random)

	return str
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if isValid(username, password) {
		session, err := store.Get(r, "server.com_"+username)
		if err != nil {
			fmt.Println("get session from session store error:", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

		// 设置session数据
		random := randomStr()
		usernameB64 := base64Encode(username + "-" + random)
		session.Values["random"] = random
		session.Save(r, w)

		// 设置cookie
		cookie := http.Cookie{Name: "server.com-session", Value: usernameB64}
		http.SetCookie(w, &cookie)

		// 登录成功,跳转到calc页面
		http.Redirect(w, r, "/calc", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized) // 401
	}
}

func calcAdd(w http.ResponseWriter, r *http.Request) {
	// 1. 获取Cookie中的Session
	cookie, err := r.Cookie("server.com-session")
	if err != nil {
		http.Error(w, "找不到cookie，请重新登录", 401)
		return
	}
	fmt.Printf("found cookie: %#v\n", cookie)

	// 2. 获取Session对象
	usernameB64 := cookie.Value
	usernameWithRandom := base64Decode(usernameB64)

	ss := strings.Split(usernameWithRandom, "-")
	username := ss[0]
	random := ss[1]
	session, err := store.Get(r, "server.com_"+username)
	if err != nil {
		http.Error(w, "找不到session, 请重新登录", 401)
		return
	}

	randomInSs := session.Values["random"]
	if random != randomInSs {
		http.Error(w, "session中信息不匹配, 请重新登录", 401)
		return
	}

	// 3. 转换为整型参数
	a, err := strconv.Atoi(r.FormValue("a"))
	if err != nil {
		http.Error(w, "参数错误", 400)
		return
	}

	b, err := strconv.Atoi(r.FormValue("b"))
	if err != nil {
		http.Error(w, "参数错误", 400)
		return
	}

	// 4. 计算并返回结果
	result := a + b
	w.Write([]byte(fmt.Sprintf("%d", result)))
}

func calc(w http.ResponseWriter, r *http.Request) {
	// 加载calc页面HTML
	calcHTML, err := os.ReadFile("./calc.html")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// 返回calc页面
	w.Header().Set("Content-Type", "text/html")
	w.Write(calcHTML)
}

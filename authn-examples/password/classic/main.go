package main

import "net/http"

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if isValidUser(username, password) {
		w.Write([]byte("Welcome!"))
		return
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized) // 401
}

var credentials = map[string]string{
	"admin": "123456",
}

func isValidUser(username, password string) bool {
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

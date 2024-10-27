// main.go
package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	ID          []byte
	Name        string
	DisplayName string
	Credentials []webauthn.Credential
}

// Implement webauthn.User interface
func (u *User) WebAuthnID() []byte {
	return u.ID
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u *User) WebAuthnIcon() string {
	return ""
}

type UserDB struct {
	sync.RWMutex
	users map[string]*User
}

func NewUserDB() *UserDB {
	return &UserDB{
		users: make(map[string]*User),
	}
}

// 用于存储会话数据
var sessionStore = make(map[string]*webauthn.SessionData)

var (
	webAuthn *webauthn.WebAuthn
	userDB   *UserDB
)

func init() {
	var err error
	webAuthn, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "Passkey Demo",                    // Relying Party Display Name
		RPID:          "localhost",                       // Relying Party ID
		RPOrigins:     []string{"http://localhost:8080"}, //允许的源
	})
	if err != nil {
		log.Fatal(err)
	}
	userDB = NewUserDB() // 初始化内存用户数据库
}

func main() {
	// 静态文件服务
	http.Handle("/", http.FileServer(http.Dir("static")))

	// API 路由
	http.HandleFunc("/api/register/begin", handleBeginRegistration)
	http.HandleFunc("/api/register/finish", handleFinishRegistration)
	http.HandleFunc("/api/login/begin", handleBeginLogin)
	http.HandleFunc("/api/login/finish", handleFinishLogin)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 辅助函数：生成会话ID并存储会话数据
func storeSession(sessionData *webauthn.SessionData) string {
	sessionID := generateSessionID()
	sessionStore[sessionID] = sessionData
	return sessionID
}

// 辅助函数：生成简单的会话ID
func generateSessionID() string {
	// 在实际应用中应使用更安全的方法生成会话ID
	b := make([]byte, 32)
	return base64.URLEncoding.EncodeToString(b)
}

// 辅助函数：获取会话数据
func getSession(sessionID string) (*webauthn.SessionData, bool) {
	session, ok := sessionStore[sessionID]
	return session, ok
}

func handleBeginRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userDB.Lock()
	if _, exists := userDB.users[data.Username]; exists {
		userDB.Unlock()
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	user := &User{
		ID:          []byte(data.Username),
		Name:        data.Username,
		DisplayName: data.Username,
	}
	userDB.users[data.Username] = user
	userDB.Unlock()

	options, sessionData, err := webAuthn.BeginRegistration(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 存储会话数据
	sessionID := storeSession(sessionData)
	http.SetCookie(w, &http.Cookie{
		Name:     "registration_session",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   300, // 5分钟过期
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(options)
}

func handleFinishRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取会话数据
	cookie, err := r.Cookie("registration_session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusBadRequest)
		return
	}

	sessionData, ok := getSession(cookie.Value)
	if !ok {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

	username := string(sessionData.UserID)
	user := userDB.users[username]

	credential, err := webAuthn.FinishRegistration(user, *sessionData, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDB.Lock()
	user.Credentials = append(user.Credentials, *credential)
	userDB.Unlock()

	// 清理会话数据
	delete(sessionStore, cookie.Value)

	w.Write([]byte("Registration success"))
}

func handleBeginLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, ok := userDB.users[data.Username]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	options, sessionData, err := webAuthn.BeginLogin(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 存储会话数据
	sessionID := storeSession(sessionData)
	http.SetCookie(w, &http.Cookie{
		Name:     "login_session",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   300, // 5分钟过期
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(options)
}

func handleFinishLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("login_session")
	if err != nil {
		http.Error(w, "Session not found", http.StatusBadRequest)
		return
	}

	sessionData, ok := getSession(cookie.Value)
	if !ok {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

	username := string(sessionData.UserID)
	user := userDB.users[username]

	_, err = webAuthn.FinishLogin(user, *sessionData, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 清理会话数据
	delete(sessionStore, cookie.Value)

	w.Write([]byte("Login success"))
}

package main

import (
	"database/sql"
	"encoding/base64"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	// 连接SQLite数据库
	var err error
	db, err = sql.Open("sqlite", "./users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建用户表
	sqltable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			hashedpass TEXT
		);
	`
	_, err = db.Exec(sqltable)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	cpassword := r.FormValue("confirm-password")

	if password != cpassword {
		http.Error(w, "password and confirmation password do not match", http.StatusBadRequest)
		return
	}

	// 注册新用户
	salt := generateSalt(16)
	hashedPassword := hashPassword(password, salt)
	stmt, err := db.Prepare("INSERT INTO users(username, hashedpass) values(?, ?)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(username, hashedPassword+":"+salt)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("signup ok!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 验证登录
	storedHashedPassword, salt := getHashedPasswordForUser(db, username)
	hashedLoginPassword := hashPassword(password, salt)
	if hashedLoginPassword == storedHashedPassword {
		w.Write([]byte("Welcome!"))
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized) // 401
	}
}

// 生成随机字符串作为盐值
func generateSalt(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 对密码进行bcrypt哈希并返回哈希值与随机盐值
func hashPassword(password, salt string) string {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}

// 从数据库获取用户哈希后的密码和盐值
func getHashedPasswordForUser(db *sql.DB, username string) (string, string) {
	var hashedPass string
	row := db.QueryRow("SELECT hashedpass FROM users WHERE username=?", username)
	if err := row.Scan(&hashedPass); err != nil {
		panic(err)
	}
	split := strings.Split(hashedPass, ":")
	return split[0], split[1]
}

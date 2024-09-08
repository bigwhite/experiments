// main.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var (
	tokens   = make(map[string]bool)
	tokensMu sync.Mutex
)

type LoginResponse struct {
	Success  bool   `json:"success"`
	Token    string `json:"token,omitempty"`
	Message  string `json:"message"`
	Redirect string `json:"redirect,omitempty"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/protected", protectedHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	response := LoginResponse{}

	if username == "admin" && password == "password" {
		token := uuid.New().String()

		tokensMu.Lock()
		tokens[token] = true
		tokensMu.Unlock()

		response.Success = true
		response.Token = token
		response.Message = "Login successful"
		response.Redirect = "/dashboard"
	} else {
		response.Success = false
		response.Message = "Login failed. Please check your credentials and try again."
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	tokensMu.Lock()
	valid := tokens[token]
	tokensMu.Unlock()

	if !valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, `<div>
        <h2>Protected Content</h2>
        <p>This is sensitive information only for authenticated users.</p>
        <p>Your token: %s</p>
    </div>`, token)
}

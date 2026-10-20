package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed templates/*.html
var tplFS embed.FS

var tpl = template.Must(template.ParseFS(tplFS, "templates/*.html"))

// ---------- 数据层：内存存储，换成 sqlc + PostgreSQL 时只需替换这里 ----------

type Plant struct {
	ID     int
	Name   string
	Health string
}

type Entry struct {
	Type  string
	Notes string
	At    time.Time
}

type Store struct {
	mu      sync.RWMutex
	plants  []*Plant
	entries map[int][]Entry
}

func NewStore() *Store {
	return &Store{
		plants: []*Plant{
			{1, "薄荷", "良好"},
			{2, "罗勒", "一般"},
			{3, "多肉", "良好"},
		},
		entries: map[int][]Entry{
			1: {{"浇水", "早上浇了一杯水", time.Now().Add(-48 * time.Hour)}},
		},
	}
}

func (s *Store) Search(q string) []*Plant {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []*Plant
	for _, p := range s.plants {
		if q == "" || strings.Contains(strings.ToLower(p.Name), strings.ToLower(q)) {
			c := *p
			out = append(out, &c)
		}
	}
	return out
}

func (s *Store) Get(id int) (*Plant, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.plants {
		if p.ID == id {
			c := *p
			return &c, true
		}
	}
	return nil, false
}

func (s *Store) SetHealth(id int, level string) (*Plant, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.plants {
		if p.ID == id {
			p.Health = level
			c := *p
			return &c, true
		}
	}
	return nil, false
}

func (s *Store) Delete(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.plants {
		if p.ID == id {
			s.plants = append(s.plants[:i], s.plants[i+1:]...)
			return
		}
	}
}

func (s *Store) AddEntry(id int, e Entry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries[id] = append([]Entry{e}, s.entries[id]...) // 新记录放最前面
}

func (s *Store) Entries(id int) []Entry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]Entry(nil), s.entries[id]...)
}

// ---------- 传输层：HTTP handler ----------

type App struct{ store *Store }

func isHX(r *http.Request) bool { return r.Header.Get("HX-Request") == "true" }

func render(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tpl.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("render %s: %v", name, err)
	}
}

func plantID(r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	return id, err == nil
}

// 首页：完整页面
func (a *App) index(w http.ResponseWriter, r *http.Request) {
	render(w, "page", map[string]any{"Q": "", "Plants": a.store.Search("")})
}

// 搜索：同一个 URL，HTMX 请求返回片段，直接访问（刷新、分享链接）返回完整页面
func (a *App) search(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	data := map[string]any{"Q": q, "Plants": a.store.Search(q)}
	w.Header().Add("Vary", "HX-Request") // 同一 URL 两种响应，告诉缓存区分对待
	if isHX(r) {
		render(w, "plant-list", data)
		return
	}
	render(w, "page", data)
}

// 详情面板（片段）
func (a *App) detail(w http.ResponseWriter, r *http.Request) {
	id, _ := plantID(r)
	p, ok := a.store.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	render(w, "detail", p)
}

// 时间线：模拟一个慢查询，配合 hx-trigger="load" 懒加载
func (a *App) timeline(w http.ResponseWriter, r *http.Request) {
	id, _ := plantID(r)
	time.Sleep(600 * time.Millisecond)
	render(w, "timeline", map[string]any{"Entries": a.store.Entries(id)})
}

// 添加事件：成功返回时间线 + OOB 提示；失败用响应头改道到 #toast
func (a *App) addEvent(w http.ResponseWriter, r *http.Request) {
	id, _ := plantID(r)
	typ := strings.TrimSpace(r.FormValue("event_type"))
	if typ == "" {
		w.Header().Set("HX-Retarget", "#toast")
		w.Header().Set("HX-Reswap", "innerHTML")
		w.WriteHeader(http.StatusUnprocessableEntity)
		render(w, "error-text", "事件类型不能为空")
		return
	}
	a.store.AddEntry(id, Entry{Type: typ, Notes: strings.TrimSpace(r.FormValue("notes")), At: time.Now()})
	render(w, "timeline", map[string]any{"Entries": a.store.Entries(id)})
	render(w, "toast", "已记录：+"+typ)
}

// 更新健康度：一次响应同时更新列表徽标、详情徽标和提示
func (a *App) setHealth(w http.ResponseWriter, r *http.Request) {
	id, _ := plantID(r)
	p, ok := a.store.SetHealth(id, r.FormValue("level"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	render(w, "health-badge-oob", p)        // 列表里的徽标
	render(w, "health-badge-detail-oob", p) // 详情里的徽标
	render(w, "toast", p.Name+" 的健康度已更新为 "+p.Health)
}

func (a *App) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := plantID(r)
	a.store.Delete(id)
	// 主体为空：目标 <li> 会被整个移除；同时用 OOB 更新提示
	render(w, "toast", "已删除")
}

// ---------- 登录过期演示：HX-Redirect ----------

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	if isHX(r) {
		w.Header().Set("HX-Redirect", url) // 普通 302 会被 HTMX 当成“要塞进目标里的页面”
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (a *App) secure(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("session"); err != nil {
		redirect(w, r, "/login")
		return
	}
	render(w, "toast-inline", "🔓 你已登录，这是受保护的内容")
}

func (a *App) loginPage(w http.ResponseWriter, r *http.Request) { render(w, "login", nil) }

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: "demo", Path: "/", HttpOnly: true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session", Path: "/", MaxAge: -1})
	redirect(w, r, "/")
}

func main() {
	a := &App{store: NewStore()}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.index)
	mux.HandleFunc("GET /plants", a.search)
	mux.HandleFunc("GET /plants/{id}", a.detail)
	mux.HandleFunc("GET /plants/{id}/timeline", a.timeline)
	mux.HandleFunc("POST /plants/{id}/events", a.addEvent)
	mux.HandleFunc("POST /plants/{id}/health", a.setHealth)
	mux.HandleFunc("DELETE /plants/{id}", a.delete)
	mux.HandleFunc("GET /secure", a.secure)
	mux.HandleFunc("GET /login", a.loginPage)
	mux.HandleFunc("POST /login", a.login)
	mux.HandleFunc("POST /logout", a.logout)

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

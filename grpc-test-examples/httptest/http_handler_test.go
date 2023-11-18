package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/plain")

	// 根据请求方法进行不同的处理
	switch r.Method {
	case http.MethodGet:
		// 处理GET请求
		fmt.Fprint(w, "Hello, World!")
	case http.MethodPost:
		// 处理POST请求
		// 获取请求正文
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 进行处理并返回响应
		response := processPostRequest(body)
		fmt.Fprint(w, response)
	default:
		// 处理其他请求方法
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func processPostRequest(body []byte) string {
	// 假设这里对请求正文进行了某种处理，并生成了响应内容
	// 这里只是一个示例，实际处理逻辑需要根据具体业务需求进行编写

	// 将请求正文解析为结构体
	var requestPayload struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(body, &requestPayload); err != nil {
		return "Error parsing request body"
	}

	// 根据请求正文中的内容生成响应
	response := fmt.Sprintf("Hello, %s!", requestPayload.Name)

	return response
}

func TestMyHandler(t *testing.T) {
	// 创建一个ResponseRecorder来记录Handler的响应
	rr := httptest.NewRecorder()

	// 创建一个模拟的HTTP请求，可以指定请求的方法、路径、正文等
	req, err := http.NewRequest("GET", "/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 调用被测试的Handler函数，传入ResponseRecorder和Request对象
	// 这里假设被测试的Handler函数为myHandler
	myHandler(rr, req)

	// 检查响应状态码和内容
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200; got %d", rr.Code)
	}
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("Expected body to be %q; got %q", expected, rr.Body.String())
	}
}

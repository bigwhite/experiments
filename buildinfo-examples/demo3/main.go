package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	_ "github.com/gin-gonic/gin"
)

// DepInfo 定义返回给前端的依赖信息结构
type DepInfo struct {
	Path    string `json:"path"`    // 依赖包路径
	Version string `json:"version"` // 依赖版本
	Sum     string `json:"sum"`     // 校验和
}

// BuildInfoResponse 完整的构建信息响应
type BuildInfoResponse struct {
	GoVersion string    `json:"go_version"`
	MainMod   string    `json:"main_mod"`
	Deps      []DepInfo `json:"deps"`
}

func depsHandler(w http.ResponseWriter, r *http.Request) {
	// 读取构建信息
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "无法获取构建信息，请确保使用 Go Modules 构建", http.StatusInternalServerError)
		return
	}

	resp := BuildInfoResponse{
		GoVersion: info.GoVersion,
		MainMod:   info.Main.Path,
		Deps:      make([]DepInfo, 0, len(info.Deps)),
	}

	// 遍历依赖树
	for _, d := range info.Deps {
		resp.Deps = append(resp.Deps, DepInfo{
			Path:    d.Path,
			Version: d.Version,
			Sum:     d.Sum,
		})
	}

	// 设置响应头并输出 JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("JSON编码失败: %v", err)
	}
}

func main() {
	http.HandleFunc("/debug/deps", depsHandler)

	fmt.Println("服务已启动，请访问: http://localhost:8080/debug/deps")
	// 为了演示依赖输出，你需要确保这个项目是一个 go mod 项目，并引入了一些第三方库
	// 例如：go get github.com/gin-gonic/gin
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"debug/buildinfo"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	// 解析命令行参数
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("用法: inspector <path-to-go-binary>")
		os.Exit(1)
	}

	binPath := flag.Arg(0)

	// 核心：使用 debug/buildinfo 读取文件，而不是 runtime
	info, err := buildinfo.ReadFile(binPath)
	if err != nil {
		log.Fatalf("读取二进制文件失败: %v", err)
	}

	fmt.Printf("=== 二进制文件分析: %s ===\n", binPath)
	fmt.Printf("Go 版本: \t%s\n", info.GoVersion)
	fmt.Printf("主模块路径: \t%s\n", info.Main.Path)

	// 提取 VCS (Git) 信息
	fmt.Println("\n[版本控制信息]")
	vcsInfo := make(map[string]string)
	for _, setting := range info.Settings {
		vcsInfo[setting.Key] = setting.Value
	}

	// 使用 tabwriter 对齐输出
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	if rev, ok := vcsInfo["vcs.revision"]; ok {
		fmt.Fprintf(w, "Commit Hash:\t%s\n", rev)
	}
	if time, ok := vcsInfo["vcs.time"]; ok {
		fmt.Fprintf(w, "Build Time:\t%s\n", time)
	}
	if mod, ok := vcsInfo["vcs.modified"]; ok {
		dirty := "否"
		if mod == "true" {
			dirty = "是 (包含未提交的更改!)"
		}
		fmt.Fprintf(w, "Dirty Build:\t%s\n", dirty)
	}
	w.Flush()

	// 打印部分依赖
	fmt.Printf("\n[依赖模块 (前5个)]\n")
	for i, dep := range info.Deps {
		if i >= 5 {
			fmt.Printf("... 以及其他 %d 个依赖\n", len(info.Deps)-5)
			break
		}
		fmt.Printf("- %s %s\n", dep.Path, dep.Version)
	}
}

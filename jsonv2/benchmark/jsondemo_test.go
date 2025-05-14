package main

import (
	"encoding/json"
	//"encoding/json/v2" // 使用gotip运行测试时使用这个v2包

	"os"
	"testing"
)

// 假设 swagger.json 文件已下载到当前目录，且内容为一个大型 JSON 对象
const swaggerFile = "swagger.json"

func BenchmarkUnmarshalSwagger(b *testing.B) {
	data, err := os.ReadFile(swaggerFile)
	if err != nil {
		b.Fatalf("Failed to read %s: %v", swaggerFile, err)
	}

	b.ResetTimer() // 重置计时器，忽略文件读取时间
	for i := 0; i < b.N; i++ {
		var out interface{} // 使用 interface{} 简化，实际场景应为具体类型
		err := json.Unmarshal(data, &out)
		if err != nil {
			b.Fatalf("Unmarshal failed: %v", err)
		}
	}
}

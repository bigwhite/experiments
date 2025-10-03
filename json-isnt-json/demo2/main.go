package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 包含浮点数的 JSON
	jsonData := []byte(`{"price": 0.1}`)

	// 定义一个结构体，使用 float64 来接收 price 字段
	var product struct {
		Price float64 `json:"price"`
	}

	// 反序列化
	if err := json.Unmarshal(jsonData, &product); err != nil {
		panic(err)
	}

	// 单独打印时，浮点数通常会以最短、最精确的十进制形式显示
	fmt.Println("Parsed price:", product.Price)

	// 当进行算术运算时，其底层的二进制不精确性就会暴露出来
	result := product.Price + 0.2
	fmt.Println("product.Price + 0.2 =", result)

	// 为了对比，直接在 Go 中进行浮点数运算
	fmt.Println("0.1 + 0.2 directly in Go =", float64(0.1)+float64(0.2))
}

package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func main() {
	// 创建一个新的Vault客户端
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println("无法创建Vault客户端:", err)
		return
	}

	// 设置Vault服务器的地址
	client.SetAddress("http://localhost:8200/")

	// 设置Vault的访问令牌（如果需要认证）
	client.SetToken("hvs.9QOJsa7zlwHO8ieW15CXXoOp")

	// 设置要写入的机密信息
	secretData := map[string]interface{}{
		"foo": "bar",
	}

	kv2 := client.KVv2("secret") // mount "secret"

	// 写入机密信息到Vault的secret/data/{key}路径下
	key := "hello"
	_, err = kv2.Put(context.Background(), key, secretData)
	if err != nil {
		fmt.Println("无法写入机密信息:", err)
		return
	}

	// 读取Vault的secret/data/{key}路径下的机密信息
	secret, err := kv2.Get(context.Background(), key)
	if err != nil {
		fmt.Println("无法读取机密信息:", err)
		return
	}

	// 打印读取到的值
	fmt.Println("读取到的值:", secret.Data)
}

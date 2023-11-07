package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/vault/api"
)

func clientAuth(vaultAddr, user, pass string) (*api.Secret, error) {
	payload := fmt.Sprintf(`{"password": "%s"}`, pass)

	req, err := http.NewRequest("POST", vaultAddr+"/v1/auth/userpass/login/"+user, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	return api.ParseSecret(bytes.NewReader(body))
}

func main() {
	vaultAddr := "http://localhost:8200"
	user := "tonybai"
	pass := "ilovegolang"

	// client登录vault服务器获取临时访问令牌
	secret, err := clientAuth(vaultAddr, user, pass)
	if err != nil {
		fmt.Printf("登录Vault失败: %v\n", err)
		return
	}
	fmt.Printf("登录Vault成功\n")
	//fmt.Printf("%#v\n", *secret.Auth)

	// 创建Vault API客户端
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("无法创建Vault客户端: %v\n", err)
		return
	}

	// 设置 Vault 地址
	client.SetAddress("http://localhost:8200")
	token := secret.Auth.ClientToken

	// 设置临时访问令牌
	client.SetToken(token)

	kv2 := client.KVv2("secret") // mount "secret"
	// 读取Vault的secret/data/{key}路径下的机密信息
	data, err := kv2.Get(context.Background(), "hello")
	if err != nil {
		fmt.Println("无法读取机密信息:", err)
		return
	}

	// 打印读取到的值
	fmt.Println("读取到的值:", data.Data)
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/userpass"
)

func main() {
	user := "tonybai"
	pass := "ilovegolang"

	// 创建Vault API客户端
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("无法创建Vault客户端: %v\n", err)
		return
	}
	// 设置 Vault 地址
	client.SetAddress("http://localhost:8200")

	// client登录vault服务器获取临时访问令牌
	userpassAuth, err := auth.NewUserpassAuth(user, &auth.Password{FromString: pass})
	if err != nil {
		fmt.Errorf("无法初始化userpass auth method: %w", err)
		return
	}

	secret, err := client.Auth().Login(context.Background(), userpassAuth)
	if err != nil {
		fmt.Errorf("登录Vault失败: %w", err)
		return
	}
	if secret == nil {
		fmt.Printf("登录后没有secret信息返回: %v\n", err)
		return
	}
	fmt.Printf("登录Vault成功\n")

	token := secret.Auth.ClientToken

	// 设置临时访问令牌
	client.SetToken(token)

	// 设置renewel watcher
	watcher, err := client.NewLifetimeWatcher(&api.LifetimeWatcherInput{
		Secret: secret,
	})
	go watcher.Start()
	defer watcher.Stop()

	kv2 := client.KVv2("secret") // mount "secret"
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case err := <-watcher.DoneCh():
			if err != nil {
				log.Printf("Failed to renew token: %v. Re-attempting login.", err)
				return
			}

			// This occurs once the token has reached max TTL.
			log.Printf("Token can no longer be renewed. Re-attempting login.")
			return

		case renewal := <-watcher.RenewCh():
			// Renewal is now over
			log.Printf("Successfully renewed: %#v", renewal)

		case <-ticker.C:
			// 每个一秒读取一次Vault的secret/data/{key}路径下的机密信息
			data, err := kv2.Get(context.Background(), "hello")
			if err != nil {
				fmt.Println("无法读取机密信息:", err)
				continue
			}
			// 打印读取到的值
			log.Println("读取到的值:", data.Data)
		}
	}
}

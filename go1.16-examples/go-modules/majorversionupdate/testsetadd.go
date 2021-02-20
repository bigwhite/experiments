package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	c := cli.SAdd("whitelist", "Obama") // 向 blacklist 中添加元素
	fmt.Println(c.Result())

}

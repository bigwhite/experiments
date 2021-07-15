package main

import (
	"fmt"

	"github.com/bigwhite/demo3/pkg/pkg1"
	_ "github.com/go-redis/redis"
	_ "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("try to LoadPlugin...")
	err := pkg1.LoadPlugin("../demo3-plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadPlugin error:", err)
		return
	}
	fmt.Println("LoadPlugin ok")
}

package main

import (
	"fmt"

	"github.com/bigwhite/demo1/pkg/pkg1"
)

func main() {
	fmt.Println("try to LoadAndInvokeSomethingFromPlugin...")
	err := pkg1.LoadAndInvokeSomethingFromPlugin("../demo5-plugins/plugin1.so.1")
	if err != nil {
		fmt.Println("LoadAndInvokeSomethingFromPlugin error:", err)
		return
	}
	fmt.Println("LoadAndInvokeSomethingFromPlugin ok")
}

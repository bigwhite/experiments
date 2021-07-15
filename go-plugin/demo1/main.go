package main

import (
	"fmt"

	"github.com/bigwhite/demo1/pkg/pkg1"
)

func main() {
	fmt.Println("try to LoadAndInvokeSomethingFromPlugin...")
	err := pkg1.LoadAndInvokeSomethingFromPlugin("../demo1-plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadAndInvokeSomethingFromPlugin error:", err)
		return
	}
	fmt.Println("LoadAndInvokeSomethingFromPlugin ok")
}

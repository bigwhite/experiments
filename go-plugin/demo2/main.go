package main

import (
	"fmt"

	_ "github.com/bigwhite/common"
	"github.com/bigwhite/demo2/pkg/pkg1"
)

func main() {
	fmt.Println("try to LoadPlugin...")
	err := pkg1.LoadPlugin("../demo2-plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadPlugin error:", err)
		return
	}
	fmt.Println("LoadPlugin ok")
	err = pkg1.LoadPlugin("../demo2-plugins/plugin1.so")
	if err != nil {
		fmt.Println("Re-LoadPlugin error:", err)
		return
	}
	fmt.Println("Re-LoadPlugin ok")
}

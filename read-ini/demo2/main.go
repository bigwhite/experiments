package main

import (
	"github.com/bigwhite/readini/pkg/config"
	"github.com/bigwhite/readini/pkg/pkg1"
)

func main() {
	err := config.InitFromFile("conf/demo.ini")
	if err != nil {
		panic(err)
	}
	pkg1.Foo()
}

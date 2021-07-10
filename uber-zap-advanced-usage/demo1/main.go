package main

import (
	"github.com/bigwhite/zap-usage/pkg/log"
)

func main() {
	//defer log.Sync()
	log.Info("demo1:", log.String("app", "start ok"),
		log.Int("major version", 2))
	//pkg1.Foo()
}

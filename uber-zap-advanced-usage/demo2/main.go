package main

import "github.com/bigwhite/zap-usage/pkg/log"

func main() {
	log.Info("demo1:", log.String("app", "start ok"))
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct {
	weaver.Implements[weaver.Main]
}

// serve is called by weaver.Run and contains the body of the application.
func serve(context.Context, *app) error {
	fmt.Println("Hello, World")
	return nil
}

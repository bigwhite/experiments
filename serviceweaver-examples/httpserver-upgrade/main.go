package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	reverser  weaver.Ref[Reverser]
	converter weaver.Ref[Converter]
	lis       weaver.Listener
}

const version = "v2"

func serve(ctx context.Context, app *app) error {
	// The lis listener will listen on a random port chosen by the operating
	// system. This behavior can be changed in the config file.
	fmt.Printf("http %s listener available on %v\n", version, app.lis)

	// Serve the /reverse endpoint.
	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("reverse %s: recv a request\n", version)
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		reversed, err := app.reverser.Get().Reverse(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "after reversing-%s, name is %s\n", version, reversed)
	})
	// Serve the /convert endpoint.
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("convert %s: recv a request\n", version)
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		converted, err := app.converter.Get().ToUpper(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "after converting-%s, name is %s\n", version, converted)
	})
	return http.Serve(app.lis, nil)
}

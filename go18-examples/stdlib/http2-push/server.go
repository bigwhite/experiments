package main

import (
	"fmt"
	"log"
	"net/http"
)

const mainJS = `document.write('Hello World!');`

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			// If it's a HTTP/2 Server.
			// Push is supported. Try pushing rather than waiting for the browser.
			if err := pusher.Push("/static/img/gopherizeme.png", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, `<html>
<head>
<title>Hello Go 1.8</title>
</head>
<body>
	<img src="/static/img/gopherizeme.png"></img>
</body>
</html>
`)
	})
	log.Fatal(http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", nil))
}

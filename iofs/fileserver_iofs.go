package main

import (
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.FS(os.DirFS("./"))))
}

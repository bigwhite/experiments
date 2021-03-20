package main

import (
	"net/http"

	gfs "github.com/bigwhite/testiofs/gofilefs"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.FS(gfs.GoFilesFS("./"))))
}

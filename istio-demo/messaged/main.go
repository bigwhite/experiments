package main

import (
	"net/http"
        "log"
)

func msgHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	log.Println("Msgd is serving the request...")

	r.ParseForm()
	m := r.FormValue("msg")

	log.Println("Msgd recv msg ok, msg=", m)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", msgHandler)

	http.ListenAndServe("0.0.0.0:9997", mux)
}

package main

import (
	"fmt"
	"net/http"
	"flag"
)

var (
	messagesvc string
	endpoint   = "send"
)

func sendMsg(msg string) error {
	url := `http://` + messagesvc + "/" + endpoint

	resp, err := http.Get(url + "?msg=" + msg)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println("service_b:v0.1 is serving the request...")

	r.ParseForm()
	m := r.FormValue("msg")

	err := sendMsg(m)
	if err != nil {
		fmt.Println("service_b:v0.1 send msg error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("service_b:v0.1 send msg ok, msg=", m)
}

func init() {
	flag.StringVar(&messagesvc, "msgsvc", "", "the external service address")
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/notify", notifyHandler)

	http.ListenAndServe("0.0.0.0:8080", mux)
}

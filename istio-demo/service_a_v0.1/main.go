package main

import (
	"fmt"
	"net/http"
)

var (
	notifysvc = "http://svcb"
	endpoint  = "notify"
)

func notifyCustomer(msg string) error {
	url := notifysvc + "/" + endpoint

		
	resp, err := http.Get(url + "?msg=" + msg)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	resp.Body.Close()
	return nil
}

// dummy function
func handlePay() error {
	return nil
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println("service_a:v0.1 is serving the request...")

	err := handlePay()
	if err != nil {
		fmt.Println("service_a:v0.1 pay error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("service_a:v0.1 pays ok")

	err = notifyCustomer("service_a:v0.1-pays-ok")
	if err != nil {
		fmt.Println("service_a:v0.1 notify error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("service_a:v0.1 notify customer ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/pay", payHandler)

	http.ListenAndServe("0.0.0.0:8080", mux)
}

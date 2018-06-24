package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request")
	fmt.Fprintf(w, "Hello, I am svc9 for ingress-controller demo!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "rootCA.pem"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8080",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
	fmt.Println("ListenAndServeTLS exit ok")
}

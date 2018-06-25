package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/*
		pool := x509.NewCertPool()
		caCertPath := "rootCA.pem"

		caCrt, err := ioutil.ReadFile(caCertPath)
		if err != nil {
			fmt.Println("ReadFile err:", err)
			return
		}
		pool.AppendCertsFromPEM(caCrt)
	*/

	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			//RootCAs:      pool,
			Certificates:       []tls.Certificate{cliCrt},
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://svc9.tonybai.com:30092")

	if err != nil {
		// handle error
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

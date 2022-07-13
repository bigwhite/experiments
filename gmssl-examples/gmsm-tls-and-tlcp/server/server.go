package main

import (
	"fmt"
	"io/ioutil"

	tls "github.com/tjfoc/gmsm/gmtls"
	x509 "github.com/tjfoc/gmsm/x509"
)

const (
	rsaCertPath     = "certs/server-rsa-cert.pem"
	rsaKeyPath      = "certs/server-rsa-key.pem"
	sm2SignCertPath = "certs/server-gm-sign-cert.pem"
	sm2SignKeyPath  = "certs/server-gm-sign-key.pem"
	sm2EncCertPath  = "certs/server-gm-enc-cert.pem"
	sm2EncKeyPath   = "certs/server-gm-enc-key.pem"
)

func main() {
	pool := x509.NewCertPool()
	rsaCACertPath := "./certs/ca-rsa-cert.pem"
	rsaCACrt, err := ioutil.ReadFile(rsaCACertPath)
	if err != nil {
		fmt.Println("read rsa ca err:", err)
		return
	}
	gmCACertPath := "./certs/ca-gm-cert.pem"
	gmCACrt, err := ioutil.ReadFile(gmCACertPath)
	if err != nil {
		fmt.Println("read gm ca err:", err)
		return
	}
	pool.AppendCertsFromPEM(rsaCACrt)
	pool.AppendCertsFromPEM(gmCACrt)

	rsaKeypair, err := tls.LoadX509KeyPair(rsaCertPath, rsaKeyPath)
	if err != nil {
		fmt.Println("load rsa x509 keypair error:", err)
		return
	}
	sigCert, err := tls.LoadX509KeyPair(sm2SignCertPath, sm2SignKeyPath)
	if err != nil {
		fmt.Println("load x509 gm sign keypair error:", err)
		return
	}
	encCert, err := tls.LoadX509KeyPair(sm2EncCertPath, sm2EncKeyPath)
	if err != nil {
		fmt.Println("load x509 gm enc keypair error:", err)
		return
	}

	cfg, err := tls.NewBasicAutoSwitchConfig(&sigCert, &encCert, &rsaKeypair)
	if err != nil {
		fmt.Println("load basic config error:", err)
		return
	}

	cfg.MaxVersion = tls.VersionTLS12
	cfg.ClientAuth = tls.RequireAndVerifyClientCert
	//cfg.RootCAs = pool
	cfg.ClientCAs = pool

	listener, err := tls.Listen("tcp", ":18000", cfg)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}
		fmt.Println("accept connection:", conn.RemoteAddr())
		go func() {
			for {
				// echo "request"
				var b = make([]byte, 16)
				_, err := conn.Read(b)
				if err != nil {
					fmt.Println("connection read error:", err)
					conn.Close()
					return
				}

				fmt.Println(string(b))
				_, err = conn.Write(b)
				if err != nil {
					fmt.Println("connection write error:", err)
					return
				}
			}
		}()
	}
}

package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"time"
)

func main() {
	// 加载ca-cert.pem
	caCertBytes, err := os.ReadFile("ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertblock, _ := pem.Decode(caCertBytes)
	caCert, err := x509.ParseCertificate(caCertblock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// 创建TLS配置
	config := &tls.Config{
		InsecureSkipVerify: true, // trigger to call VerifyPeerCertificate

		// 设置证书验证回调函数
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			// 解析服务端返回的证书链(顺序：server-cert.pem, inter-cert.pem，inter-cert.pem's issuer...)

			var issuer *x509.Certificate
			var cert *x509.Certificate
			var err error

			if len(rawCerts) == 0 {
				return errors.New("no server certificate found")
			}

			issuer = caCert

			for i := len(rawCerts) - 1; i >= 0; i-- {
				cert, err = x509.ParseCertificate(rawCerts[i])
				if err != nil {
					return err
				}

				if !verifyCert(issuer, cert) {
					return errors.New("verifyCert failed")
				}

				issuer = cert
			}
			return nil
		},
	}

	conn, err := tls.Dial("tcp", "server.com:8443", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 每秒发送信息
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		msg := "hello, tls"
		conn.Write([]byte(msg))

		// 读取回复
		buf := make([]byte, len(msg))
		conn.Read(buf)
		log.Println(string(buf))
	}

}

// 验证cert是否是issuer的签发
func verifyCert(issuer, cert *x509.Certificate) bool {

	// 验证证书
	certPool := x509.NewCertPool()
	certPool.AddCert(issuer) // ok
	opts := x509.VerifyOptions{
		Roots: certPool,
	}
	_, err := cert.Verify(opts)
	return err == nil
}

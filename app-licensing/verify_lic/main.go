package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type License struct {
	ID             string `json:"id"`
	Vendor         string `json:"vendor"`
	IssuedTo       string `json:"issuedTo"`
	IssuedDate     string `json:"issuedDate"`
	ExpirationDate string `json:"expirationDate"`
	Product        string `json:"product"`
	Version        string `json:"version"`
	LicenseType    string `json:"licenseType"`
	MaxConnections int    `json:"maxConnections"`
}

type Signature struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

func main() {

	// 1. 加载公钥证书,提取公钥
	certData, _ := os.ReadFile("ddd-cert.pem")
	block, _ := pem.Decode(certData)
	cert, _ := x509.ParseCertificate(block.Bytes)
	pubKey := cert.PublicKey.(*rsa.PublicKey)

	// 2. 解析许可证文件
	licData, err := os.ReadFile("ddd-company.lic")
	if err != nil {
		panic(err)
	}

	var license License
	var signature Signature

	err = json.Unmarshal(licData, &struct{ License *License }{&license})
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(licData, &struct{ Signature *Signature }{&signature})
	if err != nil {
		panic(err)
	}

	// 3. 生成签名摘要
	data := []string{
		license.ID,
		license.Vendor,
		license.IssuedTo,
		license.IssuedDate,
		license.ExpirationDate,
		license.Product,
		license.Version,
		license.LicenseType,
		strconv.Itoa(license.MaxConnections),
	}
	payload := strings.Join(data, "")
	hash := sha256.Sum256([]byte(payload))

	// 4. 使用公钥验签
	signValue, _ := base64.StdEncoding.DecodeString(signature.Value)
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], signValue)
	if err != nil {
		fmt.Println("Invalid signature:", err)
	} else {
		fmt.Println("Signature verified")
	}
}

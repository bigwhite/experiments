package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"log"
	"os"
	"strconv"
	"strings"
)

// 1. 建立对应license和Signature的结构体类型
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
	keyData, _ := os.ReadFile("ddd-key.pem") // 加载私钥

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing private key")
	}

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// 2. 填充license的各个字段的值
	var license License
	license.ID = "01234567890"
	license.Vendor = "XYZ Company"
	license.IssuedTo = "DDD Company"
	license.IssuedDate = "2023-10-01T00:00:00Z"
	license.ExpirationDate = "2024-09-30T23:59:59Z"
	license.Product = "My App"
	license.Version = "1.0"
	license.LicenseType = "Enterprise"
	license.MaxConnections = 1000

	// 3. 将各个字段连接后sha256摘要
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

	// 4. 用私钥对摘要签名
	signed, _ := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash[:])

	// 5. 对签名结果base64编码
	signedB64 := base64.StdEncoding.EncodeToString(signed)

	// 6. 生成signature对象
	signature := Signature{
		Algorithm: "SHA256withRSA",
		Value:     signedB64,
	}

	// 7. 序列化为json
	fullLicense := map[string]interface{}{
		"license":   license,
		"signature": signature,
	}
	jsonData, _ := json.MarshalIndent(fullLicense, "", "  ")

	// 8. 保存为.lic文件
	os.WriteFile("ddd-company.lic", jsonData, 0644)
}

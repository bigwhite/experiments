package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	// 生成CA根证书密钥对
	caKey, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)

	// 生成CA证书模板
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Go CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// 使用模板自签名生成CA证书
	caCert, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caKey.PublicKey, caKey)
	checkError(err)

	// 生成中间CA密钥对
	interKey, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)

	// 生成中间CA证书模板
	interTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"Go Intermediate CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// 用CA证书签名生成中间CA证书
	interCert, err := x509.CreateCertificate(rand.Reader, &interTemplate, &caTemplate, &interKey.PublicKey, caKey)
	checkError(err)

	// 生成叶子证书密钥对
	leafKey, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)

	// 生成叶子证书模板,CN为server.com
	leafTemplate := x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization: []string{"Go Server"},
			CommonName:   "server.com",
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
		DNSNames:     []string{"server.com"},
		SubjectKeyId: []byte{1, 2, 3, 4},
	}

	// 用中间CA证书签名生成叶子证书
	leafCert, err := x509.CreateCertificate(rand.Reader, &leafTemplate, &interTemplate, &leafKey.PublicKey, interKey)
	checkError(err)

	// 生成server1.com叶子证书
	leafKey1, _ := rsa.GenerateKey(rand.Reader, 2048)

	leafTemplate1 := x509.Certificate{
		SerialNumber: big.NewInt(4),
		Subject: pkix.Name{
			CommonName: "server1.com",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"server1.com"},
	}

	leafCert1, _ := x509.CreateCertificate(rand.Reader, &leafTemplate1, &interTemplate, &leafKey1.PublicKey, interKey)

	// 生成server2.com叶子证书
	leafKey2, _ := rsa.GenerateKey(rand.Reader, 2048)

	leafTemplate2 := x509.Certificate{
		SerialNumber: big.NewInt(5),
		Subject: pkix.Name{
			CommonName: "server2.com",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"server2.com"},
	}

	leafCert2, _ := x509.CreateCertificate(rand.Reader, &leafTemplate2, &interTemplate, &leafKey2.PublicKey, interKey)

	// 将证书和密钥编码为PEM格式
	caCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caCert})
	caKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caKey)})

	interCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: interCert})
	interKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(interKey)})

	leafCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: leafCert})
	leafKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(leafKey)})

	leafCertPEM1 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: leafCert1})
	leafKeyPEM1 := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(leafKey1)})

	leafCertPEM2 := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: leafCert2})
	leafKeyPEM2 := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(leafKey2)})

	// 将PEM写入文件
	writeDataToFile("ca-cert.pem", caCertPEM)
	writeDataToFile("ca-key.pem", caKeyPEM)

	writeDataToFile("inter-cert.pem", interCertPEM)
	writeDataToFile("inter-key.pem", interKeyPEM)

	writeDataToFile("leaf-server-cert.pem", leafCertPEM)
	writeDataToFile("leaf-server-key.pem", leafKeyPEM)

	writeDataToFile("leaf-server1-cert.pem", leafCertPEM1)
	writeDataToFile("leaf-server1-key.pem", leafKeyPEM1)

	writeDataToFile("leaf-server2-cert.pem", leafCertPEM2)
	writeDataToFile("leaf-server2-key.pem", leafKeyPEM2)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeDataToFile(fileName string, data []byte) {
	f, err := os.Create(fileName)
	checkError(err)
	defer f.Close()

	f.Write(data)
}

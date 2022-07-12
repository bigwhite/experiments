#!/bin/bash


## RSA Certs

### CA
gmssl genpkey -out ca-rsa-key.pem -algorithm RSA -pkeyopt rsa_keygen_bits:2048
gmssl req -x509 -new -nodes -key ca-rsa-key.pem -subj "/CN=myca.com" -days 5000 -out ca-rsa-cert.pem 

### server key and cert
gmssl genpkey -out server-rsa-key.pem -algorithm RSA -pkeyopt rsa_keygen_bits:2048
gmssl req -new -key server-rsa-key.pem -subj "/CN=example.com" -out server-rsa.csr
gmssl x509 -req -in server-rsa.csr -CA ca-rsa-cert.pem -CAkey ca-rsa-key.pem -CAcreateserial -out server-rsa-cert.pem -days 5000  -extfile ./server.cnf -extensions ext
gmssl verify -CAfile ca-rsa-cert.pem server-rsa-cert.pem

### client key and cert
gmssl genpkey -out client-rsa-key.pem -algorithm RSA -pkeyopt rsa_keygen_bits:2048
gmssl req -new -key client-rsa-key.pem -subj "/CN=client1.com" -out client-rsa.csr
gmssl x509 -req -in client-rsa.csr -CA ca-rsa-cert.pem -CAkey ca-rsa-key.pem -CAcreateserial -out client-rsa-cert.pem -days 5000 -extfile ./client.cnf -extensions ext
gmssl verify -CAfile ca-rsa-cert.pem client-rsa-cert.pem






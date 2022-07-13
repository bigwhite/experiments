#!/bin/bash

# execute after gen_rsa_cert.sh

## RSA Certs

### server key and cert
gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out server-gm-key.pem
gmssl req -new -key server-gm-key.pem -subj "/CN=example.com" -out server-gm.csr
gmssl x509 -req -in server-gm.csr -CA ca-rsa-cert.pem -CAkey ca-rsa-key.pem -CAcreateserial -out server-gm-cert.pem -days 5000 -extfile ./server_auth.cnf -extensions ext

gmssl verify -CAfile ca-rsa-cert.pem server-gm-cert.pem

### client key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out client-gm-key.pem
gmssl req -new -key client-gm-key.pem -subj "/CN=client1.com" -out client-gm.csr
gmssl x509 -req -in client-gm.csr -CA ca-rsa-cert.pem -CAkey ca-rsa-key.pem -CAcreateserial -out client-gm-cert.pem -days 5000 -extfile ./client_auth.cnf -extensions ext

gmssl verify -CAfile ca-rsa-cert.pem client-gm-cert.pem



#!/bin/bash


## SM CA

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out ca-gm-key.pem
gmssl req -x509 -new -nodes -key ca-gm-key.pem -subj "/CN=myca.com" -days 5000 -out ca-gm-cert.pem


### server: sign key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out server-gm-sign-key.pem
gmssl req -new -key server-gm-sign-key.pem -subj "/CN=example.com" -out server-gm-sign.csr
gmssl x509 -req -in server-gm-sign.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out server-gm-sign-cert.pem -days 5000

gmssl verify -CAfile ca-gm-cert.pem server-gm-sign-cert.pem

### server: enc key and cer

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out server-gm-enc-key.pem
gmssl req -new -key server-gm-enc-key.pem -subj "/CN=example.com" -out server-gm-enc.csr
gmssl x509 -req -in server-gm-enc.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out server-gm-enc-cert.pem -days 5000

gmssl verify -CAfile ca-gm-cert.pem server-gm-enc-cert.pem

### client: sign key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out client-gm-sign-key.pem
gmssl req -new -key client-gm-sign-key.pem -subj "/CN=client1" -out client-gm-sign.csr
gmssl x509 -req -in client-gm-sign.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out client-gm-sign-cert.pem -days 5000

gmssl verify -CAfile ca-gm-cert.pem client-gm-sign-cert.pem

### client: enc key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out client-gm-enc-key.pem
gmssl req -new -key client-gm-enc-key.pem -subj "/CN=client1" -out client-gm-enc.csr
gmssl x509 -req -in client-gm-enc.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out client-gm-enc-cert.pem -days 5000

gmssl verify -CAfile ca-gm-cert.pem client-gm-enc-cert.pem





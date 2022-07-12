#!/bin/bash


## RSA Certs

#mkcert -key-file key.pem -cert-file cert.pem example.com
#mkcert -client -key-file client-key.pem -cert-file client-cert.pem client1

## SM CA

gmssl ecparam -genkey -name sm2p256v1 -text -out ca-gm-key.pem
gmssl req -new -sm3 -key ca-gm-key.pem -out ca-gm.csr
gmssl x509 -req -sm3 -days 5000 -in ca-gm.csr -signkey ca-gm-key.pem -out ca-gm-cert.pem

#gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out ca-gm-key.pem
#gmssl req -x509 -new -nodes -key ca-gm-key.pem -days 5000 -out ca-gm-cert.pem


### server: sign key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out gm-sign-key.pem
gmssl req -new -key gm-sign-key.pem -out gm-sign.csr

gmssl x509 -req -in gm-sign.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out gm-sign-cert.pem -days 5000

### server: enc key and cer

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out gm-enc-key.pem
gmssl req -new -key gm-enc-key.pem -out gm-enc.csr
gmssl x509 -req -in gm-enc.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out gm-enc-cert.pem -days 5000

### client: sign key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out gm-client-sign-key.pem
gmssl req -new -key gm-client-sign-key.pem -out gm-client-sign.csr
gmssl x509 -req -in gm-client-sign.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out gm-client-sign-cert.pem -days 5000

### client: enc key and cert

gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 -out gm-client-enc-key.pem
gmssl req -new -key gm-client-enc-key.pem -out gm-client-enc.csr
gmssl x509 -req -in gm-client-enc.csr -CA ca-gm-cert.pem -CAkey ca-gm-key.pem -CAcreateserial -out gm-client-enc-cert.pem -days 5000







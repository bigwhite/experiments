#!/bin/bash

go install filippo.io/mkcert@latest
mkcert -install
mkcert -key-file key.pem -cert-file cert.pem example.com
mkcert -client -key-file client-key.pem -cert-file client-cert.pem client1


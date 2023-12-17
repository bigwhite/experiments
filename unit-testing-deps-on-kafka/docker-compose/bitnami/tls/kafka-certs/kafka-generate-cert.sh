openssl genpkey -algorithm RSA -out ca-key.pem
openssl req -new -key ca-key.pem -out ca-req.csr
openssl x509 -req -in ca-req.csr -signkey ca-key.pem -out ca-cert.pem -days 3650
openssl genpkey -algorithm RSA -out server-key.pem
openssl req -new -key server-key.pem -out server-req.csr
openssl x509 -req -in server-req.csr -CA ca-cert.pem -CAkey ca-key.pem -out server-cert.pem -CAcreateserial -days 3650
cp ca-cert.pem ../kafka.truststore.pem
cp server-key.pem ../kafka.keystore.key
cp server-cert.pem  ../kafka.keystore.pem

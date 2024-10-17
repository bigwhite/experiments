.PHONY: all gen client server

all: gen server client

gen:
	protoc --go_out=./proto --go-grpc_out=./proto proto/greeting.proto

server: gen
	go build -o grpc-server1 rpcdemo/server1
	go build -o grpc-server2 rpcdemo/server2

client: gen
	go build -o grpc-client rpcdemo/client

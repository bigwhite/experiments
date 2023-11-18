
all: gen

gen:
	protoc -I ./IDL/proto mygrpc.proto --gofast_out=plugins=grpc:./mygrpc

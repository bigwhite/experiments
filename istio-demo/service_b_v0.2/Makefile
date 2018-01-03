
#
## usage:
#    make build repos=xxxx tag=v0.1
#    #

# default

repos=docker.io/bigwhite/istio-demo-svcb
tag=v0.2

all: push

build:
	go build -o svcb main.go

buildimage: build 
	docker build -t $(repos):$(tag) .

push: buildimage
	docker push $(repos):$(tag)

pull:
	docker pull $(repos):$(tag)

test:
	go test -v ./...

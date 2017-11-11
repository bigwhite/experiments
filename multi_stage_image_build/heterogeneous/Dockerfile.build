FROM golang:latest

WORKDIR /go/src
COPY httpserver.go .

RUN CGO_ENABLED=0 go build -o myhttpserver ./httpserver.go

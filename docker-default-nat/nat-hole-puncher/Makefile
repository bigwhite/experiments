all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o nat-hole-puncher main.go
image:
	docker build -t nat-hole-puncher .

all:
	go build -o myapp main.go config.go
clean:
	rm -fr myapp
install:
	cp ./myapp /usr/local/bin

install-cfg:
	mkdir -p /etc/myapp
	cp ./config.ini /etc/myapp

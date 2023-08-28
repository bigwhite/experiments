all:
	go build -o foo.a -buildmode=archive
	cp foo.a ../library/github.com/bigwhite

clean:
	rm -fr foo.a
	go clean

all:
	go build -o bar.a -buildmode=archive
	cp bar.a ../library/github.com/bigwhite

clean:
	rm -fr bar.a
	go clean

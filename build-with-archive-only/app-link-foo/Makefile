all:
	go tool compile -I /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library -o main.o main.go
	go tool link -L /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library -o main main.o

clean:
	rm -fr main.o main

build-static:
	CGO_ENABLED=0 go tool compile -I /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library -o main.o main.go
	CGO_ENABLED=0 go tool link -L /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library -o main main.o


build:
	go build -x -v -gcflags '-I /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library' -ldflags '-L /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library' -o main main.go

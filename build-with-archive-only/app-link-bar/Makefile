all:
	go tool compile -I /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library -o main.o main.go
	go tool link -L /Users/tonybai/Go/src/github.com/bigwhite/experiments/build-with-archive-only/library -o main main.o

build-with-importcfg:
	go tool compile -importcfg import.link -o main.o main.go
	go tool link -importcfg import.link -o main main.o

clean:
	rm -fr main.o main

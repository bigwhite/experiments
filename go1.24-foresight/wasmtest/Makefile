all:
	GOARCH=wasm GOOS=wasip1 gotip build -buildmode=c-shared -o x.wasm ./testprog
	gotip run w.go ./x.wasm

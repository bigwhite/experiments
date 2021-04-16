package main

import (
	"fmt"
	"os"
	"reflect"
)

func fileFD(f *os.File) int {
	file := reflect.ValueOf(f).Elem().FieldByName("file").Elem()
	pfdVal := file.FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("file descriptor is %d\n", f.Fd())
	fmt.Printf("file descriptor in reflect is %d\n", fileFD(f))
}

package main

import (
	"fmt"

	"golang.org/x/exp/xerrors"
)

func main() {
	err1 := xerrors.New("1")
	err2 := xerrors.Errorf("wrap 2: %w", err1)
	err3 := xerrors.Errorf("wrap 3: %w", err2)

	erra := xerrors.New("a")

	b := xerrors.Is(err3, err1)
	fmt.Println("err3 is err1? -> ", b)

	b = xerrors.Is(err2, err1)
	fmt.Println("err2 is err1? -> ", b)

	b = xerrors.Is(err3, err2)
	fmt.Println("err3 is err2? -> ", b)

	b = xerrors.Is(erra, err1)
	fmt.Println("erra is err1? -> ", b)
}

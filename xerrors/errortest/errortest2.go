package main

import (
	"fmt"

	"golang.org/x/exp/xerrors"
)

type MyError struct{}

func (MyError) Error() string {
	return "MyError"
}

func main() {
	err1 := MyError{}
	err2 := xerrors.Errorf("wrap 2: %w", err1)
	err3 := xerrors.Errorf("wrap 3: %w", err2)
	var err MyError

	b := xerrors.As(err3, &err)
	fmt.Println("err3 as MyError? -> ", b)
	fmt.Println("err is err1? -> ", xerrors.Is(err, err1))

	err4 := xerrors.Opaque(err3)
	b = xerrors.As(err4, &err)
	fmt.Println("err4 as MyError? -> ", b)
	b = xerrors.Is(err4, err3)
	fmt.Println("err4 is err3? -> ", b)
}

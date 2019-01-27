package main

import (
	"fmt"

	"golang.org/x/exp/xerrors"
)

func function4() error {
	return xerrors.New("original_error")
}

func function3() error {
	err := function4()
	if err != nil {
		return xerrors.Errorf("wrap3: %w", err)
	}
	return nil
}

func function2() error {
	err := function3()
	if err != nil {
		return xerrors.Errorf("wrap2: %w", err)
	}
	return nil
}

func function1() error {
	err := function2()
	if err != nil {
		return xerrors.Errorf("wrap1: %w", err)
	}
	return nil
}

func main() {
	err := function1()
	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Printf("ok\n")
}

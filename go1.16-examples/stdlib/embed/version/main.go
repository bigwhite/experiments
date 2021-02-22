package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string
)

func main() {
	fmt.Printf("Version %q\n", Version)
}

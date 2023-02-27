package main

import (
	"github.com/bigwhite/privatemodule3"
	"mycompany.com/go/privatemodule1"
	"mycompany.com/go/privatemodule2"
)

func main() {
	privatemodule3.F()
	privatemodule2.F()
	privatemodule1.F()
}

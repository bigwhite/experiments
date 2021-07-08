package pkg1

import (
	"fmt"

	"github.com/bigwhite/readini/pkg/config"
)

func Foo() {
	fmt.Printf("%#v\n", config.Config)
}

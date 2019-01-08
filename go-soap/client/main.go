package main

import (
	"fmt"

	soap "github.com/bigwhite/experiments/go-soap/pkg/myservice"
)

func main() {
	//c := soap.NewCalculatorSoap("", false, nil)
	c := soap.NewCalculatorSoap("http://localhost:8080/", false, nil)
	r, err := c.Add(&soap.Add{
		IntA: 2,
		IntB: 3,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.AddResult)

	/*
		r1, err := c.Multiply(&soap.Multiply{
			IntA: 5,
			IntB: 6,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(r1.MultiplyResult)
	*/
}

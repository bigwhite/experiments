package main

import (
	"math/rand"
)

func convertSmallInteger() interface{} {
	i := rand.Intn(256)
	var j interface{} = i
	return j
}

func main() {
	for i := 0; i < 100000000; i++ {
		convertSmallInteger()
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	input, e := strconv.Atoi(os.Args[1]) // Get an input number from the command line
	if e != nil {
		panic(e)
	}
	u := int32(input)
	r := int32(rand.Intn(10000))        // Get a random number 0 <= r < 10k
	var a [10000]int32                  // Array of 10k elements initialized to 0
	for i := int32(0); i < 10000; i++ { // 10k outer loop iterations
		temp := a[i]
		for j := int32(0); j < 100000; j++ { // 100k inner loop iterations, per outer loop iteration
			temp += j % u // Simple sum
		}
		temp += r // Add a random value to each element in array
		a[i] = temp
	}
	fmt.Println(a[r]) // Print out a single element from the array
}

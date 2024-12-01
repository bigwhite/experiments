package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	input, _ := strconv.Atoi(os.Args[1]) // Get an input number from the command line
	u := int32(input)
	r := int32(rand.Uint32() % 10000)   // Use Uint32 for faster random number generation
	var a [10000]int32                  // Array of 10k elements initialized to 0
	for i := int32(0); i < 10000; i++ { // 10k outer loop iterations
		for j := int32(0); j < 100000; j++ { // 100k inner loop iterations, per outer loop iteration
			a[i] = a[i] + j%u // Simple sum
		}
		a[i] += r // Add a random value to each element in array
	}
	z := a[r]
	fmt.Println(z) // Print out a single element from the array
}

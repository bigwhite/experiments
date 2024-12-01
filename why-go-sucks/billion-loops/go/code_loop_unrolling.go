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
		var sum int32
		// Unroll inner loop in chunks of 4 for optimization
		for j := int32(0); j < 100000; j += 4 {
			sum += j % u
			sum += (j + 1) % u
			sum += (j + 2) % u
			sum += (j + 3) % u
		}
		a[i] = sum + r // Add the accumulated sum and random value
	}

	fmt.Println(a[r]) // Print out a single element from the array
}

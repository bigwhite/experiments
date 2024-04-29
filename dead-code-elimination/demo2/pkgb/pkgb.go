package pkgb

import (
	"fmt"
	"math/rand"
)

func Zoo() int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += int(rand.Int31n(1000))
	}
	fmt.Println("Zoo in pkgb")
	return sum
}

func F1() {
	fmt.Println("F1 in pkgb")
}
func F2() {
	fmt.Println("F2 in pkgb")
}

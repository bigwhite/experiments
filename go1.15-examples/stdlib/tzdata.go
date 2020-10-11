package main

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("LoadLocation error:", err)
		return
	}
	fmt.Println("LoadLocation is:", loc)
}

package array

var a [100]int
var sl = a[:]

func arrayRangeLoop() {
	var sum int
	for _, n := range a {
		sum += n
	}
}

func pointerToArrayRangeLoop() {
	var sum int
	for _, n := range &a {
		sum += n
	}
}

func sliceRangeLoop() {
	var sum int
	for _, n := range sl {
		sum += n
	}
}

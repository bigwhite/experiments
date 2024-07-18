package pkg

func MatrixAddNonSIMD(a, b, c []float32) {
	n := len(a)
	for i := 0; i < n; i++ {
		c[i] = a[i] + b[i]
	}
}

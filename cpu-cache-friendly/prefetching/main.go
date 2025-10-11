package main

// 线性访问，预取器可以完美工作
func SumLinear(data []int) int64 {
	var sum int64
	for i := 0; i < len(data); i++ {
		sum += int64(data[i])
	}
	return sum
}

// 随机访问，预取器失效
func SumRandom(data []int, indices []int) int64 {
	var sum int64
	for _, idx := range indices {
		sum += int64(data[idx])
	}
	return sum
}

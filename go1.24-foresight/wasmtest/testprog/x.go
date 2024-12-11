package main

func init() {
	println("init function called")
}

//go:wasmexport Add
func Add(a, b int64) int64 { 
	return a+b
}

func main() {
        println("hello")
}

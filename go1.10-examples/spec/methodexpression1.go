package main

func main() {
	(*struct{ error }).Error(nil)
}

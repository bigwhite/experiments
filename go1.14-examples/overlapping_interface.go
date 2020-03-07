package foo

type I interface {
	f()
	String() string
}
type J interface {
	g()
	String() string
}

type IJ interface {
	I
	J
}

package main

type MyInterface interface{
	Foo()
}

type MyInterface1 MyInterface
type MyInterface2 = MyInterface

type MyInt int

func (i *MyInt)Foo() {

}

func main() {
	var i MyInterface = new(MyInt)
	var i1 MyInterface1 = i
	var i2 MyInterface2 = i1

	print(i, i1, i2)
}

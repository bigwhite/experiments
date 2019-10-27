package main

import "fmt"

func numberLiteralBeforeGo1_13() {
	a := 53        //十进制
	b := 0700      // 八进制
	c := 0xaabbcc  // 十六进制
	c1 := 0Xddeeff // 十六进制
	f1 := 10.24
	f2 := 1.e+0
	f3 := 31415.e-4
	fmt.Printf("a's type is %T, value is %d\n", a, a)
	fmt.Printf("b's type is %T, value is %d\n", b, b)
	fmt.Printf("c's type is %T, value is %d\n", c, c)
	fmt.Printf("c1's type is %T, value is %d\n", c1, c1)
	fmt.Printf("f1's type is %T, value is %f\n", f1, f1)
	fmt.Printf("f2's type is %T, value is %f\n", f2, f2)
	fmt.Printf("f3's type is %T, value is %f\n", f3, f3)
}

func numberLiteralInGo1_13() {
	a := 5_3_7
	b := 0o700
	b1 := 0O700
	b2 := 0_700
	b3 := 0o_700
	c := 0b111
	c1 := 0B111
	c2 := 0b_111
	f1 := 0x10.24p+3
	f2 := 0x1.Fp+0
	f3 := 0x31_415.p-4
	fmt.Printf("a's type is %T, value is %d\n", a, a)
	fmt.Printf("b's type is %T, value is %d\n", b, b)
	fmt.Printf("b1's type is %T, value is %d\n", b1,b1)
	fmt.Printf("b2's type is %T, value is %d\n", b2, b2)
	fmt.Printf("b3's type is %T, value is %d\n", b3, b3)
	fmt.Printf("c's type is %T, value is %d\n",  c, c)
	fmt.Printf("c1's type is %T, value is %d\n", c1,c1)
	fmt.Printf("c2's type is %T, value is %d\n", c2,c2)
	fmt.Printf("f1's type is %T, value is %f\n", f1, f1)
	fmt.Printf("f2's type is %T, value is %f\n", f2, f2)
	fmt.Printf("f3's type is %T, value is %f\n", f3, f3)

}

func main() {
	fmt.Println("number literals before go 1.13:")
	numberLiteralBeforeGo1_13()
	fmt.Println("number literals in go 1.13:")
	numberLiteralInGo1_13()
}

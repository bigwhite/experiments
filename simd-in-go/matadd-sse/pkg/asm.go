//go:build ignore
// +build ignore

package main

import (
	"github.com/mmcloughlin/avo/attr"
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {
	TEXT("MatrixAddSIMD", attr.NOSPLIT, "func(a, b, c []float32)")
	a := Mem{Base: Load(Param("a").Base(), GP64())}
	b := Mem{Base: Load(Param("b").Base(), GP64())}
	c := Mem{Base: Load(Param("c").Base(), GP64())}
	n := Load(Param("a").Len(), GP64())

	X0 := XMM()
	X1 := XMM()

	Label("loop")
	CMPQ(n, U32(4))
	JL(LabelRef("done"))

	MOVUPS(a.Offset(0), X0)
	MOVUPS(b.Offset(0), X1)
	ADDPS(X1, X0)
	MOVUPS(X0, c.Offset(0))

	ADDQ(U32(16), a.Base)
	ADDQ(U32(16), b.Base)
	ADDQ(U32(16), c.Base)
	SUBQ(U32(4), n)
	JMP(LabelRef("loop"))

	Label("done")
	RET()

	Generate()
}

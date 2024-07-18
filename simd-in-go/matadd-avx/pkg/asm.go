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

	Y0 := YMM()
	Y1 := YMM()

	Label("loop")
	CMPQ(n, U32(8))
	JL(LabelRef("done"))

	VMOVUPS(a.Offset(0), Y0)
	VMOVUPS(b.Offset(0), Y1)
	VADDPS(Y1, Y0, Y0)
	VMOVUPS(Y0, c.Offset(0))

	ADDQ(U32(32), a.Base)
	ADDQ(U32(32), b.Base)
	ADDQ(U32(32), c.Base)
	SUBQ(U32(8), n)
	JMP(LabelRef("loop"))

	Label("done")
	RET()

	Generate()
}

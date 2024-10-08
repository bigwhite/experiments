// Code generated by command: go run asm.go -out add.s -stubs stub.go. DO NOT EDIT.

#include "textflag.h"

// func MatrixAddSIMD(a []float32, b []float32, c []float32)
// Requires: AVX
TEXT ·MatrixAddSIMD(SB), NOSPLIT, $0-72
	MOVQ a_base+0(FP), AX
	MOVQ b_base+24(FP), CX
	MOVQ c_base+48(FP), DX
	MOVQ a_len+8(FP), BX

loop:
	CMPQ    BX, $0x00000008
	JL      done
	VMOVUPS (AX), Y0
	VMOVUPS (CX), Y1
	VADDPS  Y1, Y0, Y0
	VMOVUPS Y0, (DX)
	ADDQ    $0x00000020, AX
	ADDQ    $0x00000020, CX
	ADDQ    $0x00000020, DX
	SUBQ    $0x00000008, BX
	JMP     loop

done:
	RET

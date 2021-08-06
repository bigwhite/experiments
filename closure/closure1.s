"".foo STEXT size=191 args=0x8 locals=0x20 funcid=0x0
	0x0000 00000 (closure1.go:5)	TEXT	"".foo(SB), ABIInternal, $32-8
	0x0000 00000 (closure1.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (closure1.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (closure1.go:5)	PCDATA	$0, $-2
	0x000d 00013 (closure1.go:5)	JLS	181
	0x0013 00019 (closure1.go:5)	PCDATA	$0, $-1
	0x0013 00019 (closure1.go:5)	SUBQ	$32, SP
	0x0017 00023 (closure1.go:5)	MOVQ	BP, 24(SP)
	0x001c 00028 (closure1.go:5)	LEAQ	24(SP), BP
	0x0021 00033 (closure1.go:5)	FUNCDATA	$0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
	0x0021 00033 (closure1.go:5)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (closure1.go:6)	LEAQ	type.[16]int(SB), AX
	0x0028 00040 (closure1.go:6)	MOVQ	AX, (SP)
	0x002c 00044 (closure1.go:6)	PCDATA	$1, $0
	0x002c 00044 (closure1.go:6)	CALL	runtime.newobject(SB)
	0x0031 00049 (closure1.go:6)	MOVQ	8(SP), AX
	0x0036 00054 (closure1.go:6)	MOVQ	AX, ""..autotmp_7+16(SP)
	0x003b 00059 (closure1.go:6)	MOVQ	$10, (AX)
	0x0042 00066 (closure1.go:6)	MOVQ	$11, 8(AX)
	0x004a 00074 (closure1.go:6)	MOVQ	$128, 120(AX)
	0x0052 00082 (closure1.go:7)	LEAQ	type.noalg.struct { F uintptr; "".i []int }(SB), CX
	0x0059 00089 (closure1.go:7)	MOVQ	CX, (SP)
	0x005d 00093 (closure1.go:7)	PCDATA	$1, $1
	0x005d 00093 (closure1.go:7)	NOP
	0x0060 00096 (closure1.go:7)	CALL	runtime.newobject(SB)
	0x0065 00101 (closure1.go:7)	MOVQ	8(SP), AX
	0x006a 00106 (closure1.go:7)	LEAQ	"".foo.func1(SB), CX
	0x0071 00113 (closure1.go:7)	MOVQ	CX, (AX)
	0x0074 00116 (closure1.go:7)	MOVQ	$16, 16(AX)
	0x007c 00124 (closure1.go:7)	MOVQ	$16, 24(AX)
	0x0084 00132 (closure1.go:7)	PCDATA	$0, $-2
	0x0084 00132 (closure1.go:7)	CMPL	runtime.writeBarrier(SB), $0
	0x008b 00139 (closure1.go:7)	JNE	165
	0x008d 00141 (closure1.go:7)	MOVQ	""..autotmp_7+16(SP), CX
	0x0092 00146 (closure1.go:7)	MOVQ	CX, 8(AX)
	0x0096 00150 (closure1.go:7)	PCDATA	$0, $-1
	0x0096 00150 (closure1.go:7)	MOVQ	AX, "".~r0+40(SP)
	0x009b 00155 (closure1.go:7)	MOVQ	24(SP), BP
	0x00a0 00160 (closure1.go:7)	ADDQ	$32, SP
	0x00a4 00164 (closure1.go:7)	RET
	0x00a5 00165 (closure1.go:7)	PCDATA	$0, $-2
	0x00a5 00165 (closure1.go:7)	LEAQ	8(AX), DI
	0x00a9 00169 (closure1.go:7)	MOVQ	""..autotmp_7+16(SP), CX
	0x00ae 00174 (closure1.go:7)	CALL	runtime.gcWriteBarrierCX(SB)
	0x00b3 00179 (closure1.go:7)	JMP	150
	0x00b5 00181 (closure1.go:7)	NOP
	0x00b5 00181 (closure1.go:5)	PCDATA	$1, $-1
	0x00b5 00181 (closure1.go:5)	PCDATA	$0, $-2
	0x00b5 00181 (closure1.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x00ba 00186 (closure1.go:5)	PCDATA	$0, $-1
	0x00ba 00186 (closure1.go:5)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a2  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 20 48 89 6c 24 18 48 8d 6c 24  ...H.. H.l$.H.l$
	0x0020 18 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00  .H......H..$....
	0x0030 00 48 8b 44 24 08 48 89 44 24 10 48 c7 00 0a 00  .H.D$.H.D$.H....
	0x0040 00 00 48 c7 40 08 0b 00 00 00 48 c7 40 78 80 00  ..H.@.....H.@x..
	0x0050 00 00 48 8d 0d 00 00 00 00 48 89 0c 24 0f 1f 00  ..H......H..$...
	0x0060 e8 00 00 00 00 48 8b 44 24 08 48 8d 0d 00 00 00  .....H.D$.H.....
	0x0070 00 48 89 08 48 c7 40 10 10 00 00 00 48 c7 40 18  .H..H.@.....H.@.
	0x0080 10 00 00 00 83 3d 00 00 00 00 00 75 18 48 8b 4c  .....=.....u.H.L
	0x0090 24 10 48 89 48 08 48 89 44 24 28 48 8b 6c 24 18  $.H.H.H.D$(H.l$.
	0x00a0 48 83 c4 20 c3 48 8d 78 08 48 8b 4c 24 10 e8 00  H.. .H.x.H.L$...
	0x00b0 00 00 00 eb e1 e8 00 00 00 00 e9 41 ff ff ff     ...........A...
	rel 5+4 t=17 TLS+0
	rel 36+4 t=16 type.[16]int+0
	rel 45+4 t=8 runtime.newobject+0
	rel 85+4 t=16 type.noalg.struct { F uintptr; "".i []int }+0
	rel 97+4 t=8 runtime.newobject+0
	rel 109+4 t=16 "".foo.func1+0
	rel 134+4 t=16 runtime.writeBarrier+-1
	rel 175+4 t=8 runtime.gcWriteBarrierCX+0
	rel 182+4 t=8 runtime.morestack_noctxt+0
"".bar STEXT size=175 args=0x0 locals=0x58 funcid=0x0
	0x0000 00000 (closure1.go:13)	TEXT	"".bar(SB), ABIInternal, $88-0
	0x0000 00000 (closure1.go:13)	MOVQ	(TLS), CX
	0x0009 00009 (closure1.go:13)	CMPQ	SP, 16(CX)
	0x000d 00013 (closure1.go:13)	PCDATA	$0, $-2
	0x000d 00013 (closure1.go:13)	JLS	165
	0x0013 00019 (closure1.go:13)	PCDATA	$0, $-1
	0x0013 00019 (closure1.go:13)	SUBQ	$88, SP
	0x0017 00023 (closure1.go:13)	MOVQ	BP, 80(SP)
	0x001c 00028 (closure1.go:13)	LEAQ	80(SP), BP
	0x0021 00033 (closure1.go:13)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (closure1.go:13)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0021 00033 (closure1.go:13)	FUNCDATA	$2, "".bar.stkobj(SB)
	0x0021 00033 (closure1.go:14)	PCDATA	$1, $0
	0x0021 00033 (closure1.go:14)	CALL	"".foo(SB)
	0x0026 00038 (closure1.go:14)	MOVQ	(SP), DX
	0x002a 00042 (closure1.go:15)	MOVQ	(DX), AX
	0x002d 00045 (closure1.go:15)	MOVQ	$5, (SP)
	0x0035 00053 (closure1.go:15)	CALL	AX
	0x0037 00055 (closure1.go:15)	MOVQ	8(SP), AX
	0x003c 00060 (closure1.go:16)	MOVQ	AX, (SP)
	0x0040 00064 (closure1.go:16)	CALL	runtime.convT64(SB)
	0x0045 00069 (closure1.go:16)	MOVQ	8(SP), AX
	0x004a 00074 (closure1.go:16)	XORPS	X0, X0
	0x004d 00077 (closure1.go:16)	MOVUPS	X0, ""..autotmp_14+64(SP)
	0x0052 00082 (closure1.go:16)	LEAQ	type.int(SB), CX
	0x0059 00089 (closure1.go:16)	MOVQ	CX, ""..autotmp_14+64(SP)
	0x005e 00094 (closure1.go:16)	MOVQ	AX, ""..autotmp_14+72(SP)
	0x0063 00099 (<unknown line number>)	NOP
	0x0063 00099 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), AX
	0x006a 00106 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), CX
	0x0071 00113 ($GOROOT/src/fmt/print.go:274)	MOVQ	CX, (SP)
	0x0075 00117 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 8(SP)
	0x007a 00122 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_14+64(SP), AX
	0x007f 00127 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 16(SP)
	0x0084 00132 ($GOROOT/src/fmt/print.go:274)	MOVQ	$1, 24(SP)
	0x008d 00141 ($GOROOT/src/fmt/print.go:274)	MOVQ	$1, 32(SP)
	0x0096 00150 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x009b 00155 (closure1.go:16)	MOVQ	80(SP), BP
	0x00a0 00160 (closure1.go:16)	ADDQ	$88, SP
	0x00a4 00164 (closure1.go:16)	RET
	0x00a5 00165 (closure1.go:16)	NOP
	0x00a5 00165 (closure1.go:13)	PCDATA	$1, $-1
	0x00a5 00165 (closure1.go:13)	PCDATA	$0, $-2
	0x00a5 00165 (closure1.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x00aa 00170 (closure1.go:13)	PCDATA	$0, $-1
	0x00aa 00170 (closure1.go:13)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 92  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 58 48 89 6c 24 50 48 8d 6c 24  ...H..XH.l$PH.l$
	0x0020 50 e8 00 00 00 00 48 8b 14 24 48 8b 02 48 c7 04  P.....H..$H..H..
	0x0030 24 05 00 00 00 ff d0 48 8b 44 24 08 48 89 04 24  $......H.D$.H..$
	0x0040 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0 0f 11 44  .....H.D$..W...D
	0x0050 24 40 48 8d 0d 00 00 00 00 48 89 4c 24 40 48 89  $@H......H.L$@H.
	0x0060 44 24 48 48 8b 05 00 00 00 00 48 8d 0d 00 00 00  D$HH......H.....
	0x0070 00 48 89 0c 24 48 89 44 24 08 48 8d 44 24 40 48  .H..$H.D$.H.D$@H
	0x0080 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44  .D$.H.D$.....H.D
	0x0090 24 20 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 50  $ .........H.l$P
	0x00a0 48 83 c4 58 c3 e8 00 00 00 00 e9 51 ff ff ff     H..X.......Q...
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.*os.File+0
	rel 5+4 t=17 TLS+0
	rel 34+4 t=8 "".foo+0
	rel 53+0 t=11 +0
	rel 65+4 t=8 runtime.convT64+0
	rel 85+4 t=16 type.int+0
	rel 102+4 t=16 os.Stdout+0
	rel 109+4 t=16 go.itab.*os.File,io.Writer+0
	rel 151+4 t=8 fmt.Fprintln+0
	rel 166+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=185 args=0x0 locals=0x58 funcid=0x0
	0x0000 00000 (closure1.go:19)	TEXT	"".main(SB), ABIInternal, $88-0
	0x0000 00000 (closure1.go:19)	MOVQ	(TLS), CX
	0x0009 00009 (closure1.go:19)	CMPQ	SP, 16(CX)
	0x000d 00013 (closure1.go:19)	PCDATA	$0, $-2
	0x000d 00013 (closure1.go:19)	JLS	175
	0x0013 00019 (closure1.go:19)	PCDATA	$0, $-1
	0x0013 00019 (closure1.go:19)	SUBQ	$88, SP
	0x0017 00023 (closure1.go:19)	MOVQ	BP, 80(SP)
	0x001c 00028 (closure1.go:19)	LEAQ	80(SP), BP
	0x0021 00033 (closure1.go:19)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (closure1.go:19)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0021 00033 (closure1.go:19)	FUNCDATA	$2, "".main.stkobj(SB)
	0x0021 00033 (closure1.go:20)	PCDATA	$1, $0
	0x0021 00033 (closure1.go:20)	CALL	"".bar(SB)
	0x0026 00038 (closure1.go:21)	CALL	"".foo(SB)
	0x002b 00043 (closure1.go:21)	MOVQ	(SP), DX
	0x002f 00047 (closure1.go:22)	MOVQ	(DX), AX
	0x0032 00050 (closure1.go:22)	MOVQ	$6, (SP)
	0x003a 00058 (closure1.go:22)	CALL	AX
	0x003c 00060 (closure1.go:22)	MOVQ	8(SP), AX
	0x0041 00065 (closure1.go:23)	MOVQ	AX, (SP)
	0x0045 00069 (closure1.go:23)	CALL	runtime.convT64(SB)
	0x004a 00074 (closure1.go:23)	MOVQ	8(SP), AX
	0x004f 00079 (closure1.go:23)	XORPS	X0, X0
	0x0052 00082 (closure1.go:23)	MOVUPS	X0, ""..autotmp_14+64(SP)
	0x0057 00087 (closure1.go:23)	LEAQ	type.int(SB), CX
	0x005e 00094 (closure1.go:23)	MOVQ	CX, ""..autotmp_14+64(SP)
	0x0063 00099 (closure1.go:23)	MOVQ	AX, ""..autotmp_14+72(SP)
	0x0068 00104 (<unknown line number>)	NOP
	0x0068 00104 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), AX
	0x006f 00111 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), CX
	0x0076 00118 ($GOROOT/src/fmt/print.go:274)	MOVQ	CX, (SP)
	0x007a 00122 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 8(SP)
	0x007f 00127 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_14+64(SP), AX
	0x0084 00132 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 16(SP)
	0x0089 00137 ($GOROOT/src/fmt/print.go:274)	MOVQ	$1, 24(SP)
	0x0092 00146 ($GOROOT/src/fmt/print.go:274)	MOVQ	$1, 32(SP)
	0x009b 00155 ($GOROOT/src/fmt/print.go:274)	NOP
	0x00a0 00160 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00a5 00165 (closure1.go:23)	MOVQ	80(SP), BP
	0x00aa 00170 (closure1.go:23)	ADDQ	$88, SP
	0x00ae 00174 (closure1.go:23)	RET
	0x00af 00175 (closure1.go:23)	NOP
	0x00af 00175 (closure1.go:19)	PCDATA	$1, $-1
	0x00af 00175 (closure1.go:19)	PCDATA	$0, $-2
	0x00af 00175 (closure1.go:19)	CALL	runtime.morestack_noctxt(SB)
	0x00b4 00180 (closure1.go:19)	PCDATA	$0, $-1
	0x00b4 00180 (closure1.go:19)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 9c  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 58 48 89 6c 24 50 48 8d 6c 24  ...H..XH.l$PH.l$
	0x0020 50 e8 00 00 00 00 e8 00 00 00 00 48 8b 14 24 48  P..........H..$H
	0x0030 8b 02 48 c7 04 24 06 00 00 00 ff d0 48 8b 44 24  ..H..$......H.D$
	0x0040 08 48 89 04 24 e8 00 00 00 00 48 8b 44 24 08 0f  .H..$.....H.D$..
	0x0050 57 c0 0f 11 44 24 40 48 8d 0d 00 00 00 00 48 89  W...D$@H......H.
	0x0060 4c 24 40 48 89 44 24 48 48 8b 05 00 00 00 00 48  L$@H.D$HH......H
	0x0070 8d 0d 00 00 00 00 48 89 0c 24 48 89 44 24 08 48  ......H..$H.D$.H
	0x0080 8d 44 24 40 48 89 44 24 10 48 c7 44 24 18 01 00  .D$@H.D$.H.D$...
	0x0090 00 00 48 c7 44 24 20 01 00 00 00 0f 1f 44 00 00  ..H.D$ ......D..
	0x00a0 e8 00 00 00 00 48 8b 6c 24 50 48 83 c4 58 c3 e8  .....H.l$PH..X..
	0x00b0 00 00 00 00 e9 47 ff ff ff                       .....G...
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.*os.File+0
	rel 5+4 t=17 TLS+0
	rel 34+4 t=8 "".bar+0
	rel 39+4 t=8 "".foo+0
	rel 58+0 t=11 +0
	rel 70+4 t=8 runtime.convT64+0
	rel 90+4 t=16 type.int+0
	rel 107+4 t=16 os.Stdout+0
	rel 114+4 t=16 go.itab.*os.File,io.Writer+0
	rel 161+4 t=8 fmt.Fprintln+0
	rel 176+4 t=8 runtime.morestack_noctxt+0
"".foo.func1 STEXT nosplit size=58 args=0x10 locals=0x18 funcid=0x0
	0x0000 00000 (closure1.go:7)	TEXT	"".foo.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $24-16
	0x0000 00000 (closure1.go:7)	SUBQ	$24, SP
	0x0004 00004 (closure1.go:7)	MOVQ	BP, 16(SP)
	0x0009 00009 (closure1.go:7)	LEAQ	16(SP), BP
	0x000e 00014 (closure1.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (closure1.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (closure1.go:8)	MOVQ	16(DX), CX
	0x0012 00018 (closure1.go:8)	MOVQ	8(DX), DX
	0x0016 00022 (closure1.go:8)	TESTQ	CX, CX
	0x0019 00025 (closure1.go:8)	JLS	50
	0x001b 00027 (closure1.go:8)	MOVQ	"".n+32(SP), AX
	0x0020 00032 (closure1.go:8)	ADDQ	(DX), AX
	0x0023 00035 (closure1.go:9)	MOVQ	AX, "".~r1+40(SP)
	0x0028 00040 (closure1.go:9)	MOVQ	16(SP), BP
	0x002d 00045 (closure1.go:9)	ADDQ	$24, SP
	0x0031 00049 (closure1.go:9)	RET
	0x0032 00050 (closure1.go:8)	XORL	AX, AX
	0x0034 00052 (closure1.go:8)	PCDATA	$1, $0
	0x0034 00052 (closure1.go:8)	CALL	runtime.panicIndex(SB)
	0x0039 00057 (closure1.go:8)	XCHGL	AX, AX
	0x0000 48 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 8b  H...H.l$.H.l$.H.
	0x0010 4a 10 48 8b 52 08 48 85 c9 76 17 48 8b 44 24 20  J.H.R.H..v.H.D$ 
	0x0020 48 03 02 48 89 44 24 28 48 8b 6c 24 10 48 83 c4  H..H.D$(H.l$.H..
	0x0030 18 c3 31 c0 e8 00 00 00 00 90                    ..1.......
	rel 53+4 t=8 runtime.panicIndex+0
os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0000 00000 (<autogenerated>:1)	MOVQ	""..this+8(SP), AX
	0x0005 00005 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0008 00008 (<autogenerated>:1)	MOVQ	AX, ""..this+8(SP)
	0x000d 00013 (<autogenerated>:1)	XORPS	X0, X0
	0x0010 00016 (<autogenerated>:1)	MOVUPS	X0, "".~r0+16(SP)
	0x0015 00021 (<autogenerated>:1)	JMP	os.(*file).close(SB)
	0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
	0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
	rel 22+4 t=8 os.(*file).close+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.error+0
	rel 0+0 t=24 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*[]int- SRODATA dupok size=9
	0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 44+4 t=6 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..eqfunc128 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 80 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[16]int- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 31 36 5d 69 6e 74                 ...*[16]int
type.*[16]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 6f c1 15 1d 08 08 08 36 00 00 00 00 00 00 00 00  o......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[16]int-+0
	rel 48+8 t=1 type.[16]int+0
runtime.gcbits. SRODATA dupok size=0
type.[16]int SRODATA dupok size=72
	0x0000 80 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 24 0a 4c 21 0a 08 08 11 00 00 00 00 00 00 00 00  $.L!............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc128+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[16]int-+0
	rel 44+4 t=6 type.*[16]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
type..namedata.*struct { F uintptr; i []int }- SRODATA dupok size=33
	0x0000 00 00 1e 2a 73 74 72 75 63 74 20 7b 20 46 20 75  ...*struct { F u
	0x0010 69 6e 74 70 74 72 3b 20 69 20 5b 5d 69 6e 74 20  intptr; i []int 
	0x0020 7d                                               }
type.*struct { F uintptr; "".i []int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 aa 60 d2 32 08 08 08 36 00 00 00 00 00 00 00 00  .`.2...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i []int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; "".i []int }+0
type..namedata..F- SRODATA dupok size=5
	0x0000 00 00 02 2e 46                                   ....F
type..namedata.i- SRODATA dupok size=4
	0x0000 00 00 01 69                                      ...i
type.noalg.struct { F uintptr; "".i []int } SRODATA dupok size=128
	0x0000 20 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00   ...............
	0x0010 3a 4c 12 dc 02 08 08 19 00 00 00 00 00 00 00 00  :L..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i []int }-+0
	rel 44+4 t=6 type.*struct { F uintptr; "".i []int }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; "".i []int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.i-+0
	rel 112+8 t=1 type.[]int+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=1 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
"".bar.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
"".main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........

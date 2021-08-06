"".foo STEXT size=273 args=0x8 locals=0x30 funcid=0x0
	0x0000 00000 (closure2.go:8)	TEXT	"".foo(SB), ABIInternal, $48-8
	0x0000 00000 (closure2.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (closure2.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (closure2.go:8)	PCDATA	$0, $-2
	0x000d 00013 (closure2.go:8)	JLS	263
	0x0013 00019 (closure2.go:8)	PCDATA	$0, $-1
	0x0013 00019 (closure2.go:8)	SUBQ	$48, SP
	0x0017 00023 (closure2.go:8)	MOVQ	BP, 40(SP)
	0x001c 00028 (closure2.go:8)	LEAQ	40(SP), BP
	0x0021 00033 (closure2.go:8)	FUNCDATA	$0, gclocals·5f7ae22b544db82d5d4c812af83655e9(SB)
	0x0021 00033 (closure2.go:8)	FUNCDATA	$1, gclocals·c8a67faffb154eea2f3f8b30fe59d7e4(SB)
	0x0021 00033 (closure2.go:9)	LEAQ	type.int(SB), AX
	0x0028 00040 (closure2.go:9)	MOVQ	AX, (SP)
	0x002c 00044 (closure2.go:9)	PCDATA	$1, $0
	0x002c 00044 (closure2.go:9)	CALL	runtime.newobject(SB)
	0x0031 00049 (closure2.go:9)	MOVQ	8(SP), AX
	0x0036 00054 (closure2.go:9)	MOVQ	AX, "".&a+32(SP)
	0x003b 00059 (closure2.go:9)	MOVQ	$11, (AX)
	0x0042 00066 (closure2.go:9)	LEAQ	type.int(SB), CX
	0x0049 00073 (closure2.go:9)	MOVQ	CX, (SP)
	0x004d 00077 (closure2.go:9)	PCDATA	$1, $1
	0x004d 00077 (closure2.go:9)	CALL	runtime.newobject(SB)
	0x0052 00082 (closure2.go:9)	MOVQ	8(SP), AX
	0x0057 00087 (closure2.go:9)	MOVQ	AX, "".&b+24(SP)
	0x005c 00092 (closure2.go:9)	MOVQ	$12, (AX)
	0x0063 00099 (closure2.go:9)	LEAQ	type.int(SB), CX
	0x006a 00106 (closure2.go:9)	MOVQ	CX, (SP)
	0x006e 00110 (closure2.go:9)	PCDATA	$1, $2
	0x006e 00110 (closure2.go:9)	CALL	runtime.newobject(SB)
	0x0073 00115 (closure2.go:9)	MOVQ	8(SP), AX
	0x0078 00120 (closure2.go:9)	MOVQ	AX, "".&c+16(SP)
	0x007d 00125 (closure2.go:9)	MOVQ	$13, (AX)
	0x0084 00132 (closure2.go:10)	LEAQ	type.noalg.struct { F uintptr; "".a *int; "".b *int; "".c *int }(SB), CX
	0x008b 00139 (closure2.go:10)	MOVQ	CX, (SP)
	0x008f 00143 (closure2.go:10)	PCDATA	$1, $3
	0x008f 00143 (closure2.go:10)	CALL	runtime.newobject(SB)
	0x0094 00148 (closure2.go:10)	MOVQ	8(SP), AX
	0x0099 00153 (closure2.go:10)	LEAQ	"".foo.func1(SB), CX
	0x00a0 00160 (closure2.go:10)	MOVQ	CX, (AX)
	0x00a3 00163 (closure2.go:10)	PCDATA	$0, $-2
	0x00a3 00163 (closure2.go:10)	CMPL	runtime.writeBarrier(SB), $0
	0x00aa 00170 (closure2.go:10)	JNE	214
	0x00ac 00172 (closure2.go:10)	MOVQ	"".&a+32(SP), CX
	0x00b1 00177 (closure2.go:10)	MOVQ	CX, 8(AX)
	0x00b5 00181 (closure2.go:10)	MOVQ	"".&b+24(SP), CX
	0x00ba 00186 (closure2.go:10)	MOVQ	CX, 16(AX)
	0x00be 00190 (closure2.go:10)	MOVQ	"".&c+16(SP), CX
	0x00c3 00195 (closure2.go:10)	MOVQ	CX, 24(AX)
	0x00c7 00199 (closure2.go:10)	PCDATA	$0, $-1
	0x00c7 00199 (closure2.go:10)	MOVQ	AX, "".~r0+56(SP)
	0x00cc 00204 (closure2.go:10)	MOVQ	40(SP), BP
	0x00d1 00209 (closure2.go:10)	ADDQ	$48, SP
	0x00d5 00213 (closure2.go:10)	RET
	0x00d6 00214 (closure2.go:10)	PCDATA	$0, $-2
	0x00d6 00214 (closure2.go:10)	LEAQ	8(AX), DI
	0x00da 00218 (closure2.go:10)	MOVQ	"".&a+32(SP), CX
	0x00df 00223 (closure2.go:10)	NOP
	0x00e0 00224 (closure2.go:10)	CALL	runtime.gcWriteBarrierCX(SB)
	0x00e5 00229 (closure2.go:10)	LEAQ	16(AX), DI
	0x00e9 00233 (closure2.go:10)	MOVQ	"".&b+24(SP), CX
	0x00ee 00238 (closure2.go:10)	CALL	runtime.gcWriteBarrierCX(SB)
	0x00f3 00243 (closure2.go:10)	LEAQ	24(AX), DI
	0x00f7 00247 (closure2.go:10)	MOVQ	"".&c+16(SP), CX
	0x00fc 00252 (closure2.go:10)	NOP
	0x0100 00256 (closure2.go:10)	CALL	runtime.gcWriteBarrierCX(SB)
	0x0105 00261 (closure2.go:10)	JMP	199
	0x0107 00263 (closure2.go:10)	NOP
	0x0107 00263 (closure2.go:8)	PCDATA	$1, $-1
	0x0107 00263 (closure2.go:8)	PCDATA	$0, $-2
	0x0107 00263 (closure2.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x010c 00268 (closure2.go:8)	PCDATA	$0, $-1
	0x010c 00268 (closure2.go:8)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 f4  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
	0x0020 28 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00  (H......H..$....
	0x0030 00 48 8b 44 24 08 48 89 44 24 20 48 c7 00 0b 00  .H.D$.H.D$ H....
	0x0040 00 00 48 8d 0d 00 00 00 00 48 89 0c 24 e8 00 00  ..H......H..$...
	0x0050 00 00 48 8b 44 24 08 48 89 44 24 18 48 c7 00 0c  ..H.D$.H.D$.H...
	0x0060 00 00 00 48 8d 0d 00 00 00 00 48 89 0c 24 e8 00  ...H......H..$..
	0x0070 00 00 00 48 8b 44 24 08 48 89 44 24 10 48 c7 00  ...H.D$.H.D$.H..
	0x0080 0d 00 00 00 48 8d 0d 00 00 00 00 48 89 0c 24 e8  ....H......H..$.
	0x0090 00 00 00 00 48 8b 44 24 08 48 8d 0d 00 00 00 00  ....H.D$.H......
	0x00a0 48 89 08 83 3d 00 00 00 00 00 75 2a 48 8b 4c 24  H...=.....u*H.L$
	0x00b0 20 48 89 48 08 48 8b 4c 24 18 48 89 48 10 48 8b   H.H.H.L$.H.H.H.
	0x00c0 4c 24 10 48 89 48 18 48 89 44 24 38 48 8b 6c 24  L$.H.H.H.D$8H.l$
	0x00d0 28 48 83 c4 30 c3 48 8d 78 08 48 8b 4c 24 20 90  (H..0.H.x.H.L$ .
	0x00e0 e8 00 00 00 00 48 8d 78 10 48 8b 4c 24 18 e8 00  .....H.x.H.L$...
	0x00f0 00 00 00 48 8d 78 18 48 8b 4c 24 10 0f 1f 40 00  ...H.x.H.L$...@.
	0x0100 e8 00 00 00 00 eb c0 e8 00 00 00 00 e9 ef fe ff  ................
	0x0110 ff                                               .
	rel 5+4 t=17 TLS+0
	rel 36+4 t=16 type.int+0
	rel 45+4 t=8 runtime.newobject+0
	rel 69+4 t=16 type.int+0
	rel 78+4 t=8 runtime.newobject+0
	rel 102+4 t=16 type.int+0
	rel 111+4 t=8 runtime.newobject+0
	rel 135+4 t=16 type.noalg.struct { F uintptr; "".a *int; "".b *int; "".c *int }+0
	rel 144+4 t=8 runtime.newobject+0
	rel 156+4 t=16 "".foo.func1+0
	rel 165+4 t=16 runtime.writeBarrier+-1
	rel 225+4 t=8 runtime.gcWriteBarrierCX+0
	rel 239+4 t=8 runtime.gcWriteBarrierCX+0
	rel 257+4 t=8 runtime.gcWriteBarrierCX+0
	rel 264+4 t=8 runtime.morestack_noctxt+0
"".bar STEXT size=778 args=0x0 locals=0xe8 funcid=0x0
	0x0000 00000 (closure2.go:25)	TEXT	"".bar(SB), ABIInternal, $232-0
	0x0000 00000 (closure2.go:25)	MOVQ	(TLS), CX
	0x0009 00009 (closure2.go:25)	LEAQ	-104(SP), AX
	0x000e 00014 (closure2.go:25)	CMPQ	AX, 16(CX)
	0x0012 00018 (closure2.go:25)	PCDATA	$0, $-2
	0x0012 00018 (closure2.go:25)	JLS	768
	0x0018 00024 (closure2.go:25)	PCDATA	$0, $-1
	0x0018 00024 (closure2.go:25)	SUBQ	$232, SP
	0x001f 00031 (closure2.go:25)	MOVQ	BP, 224(SP)
	0x0027 00039 (closure2.go:25)	LEAQ	224(SP), BP
	0x002f 00047 (closure2.go:25)	FUNCDATA	$0, gclocals·0ce64bbc7cfa5ef04d41c861de81a3d7(SB)
	0x002f 00047 (closure2.go:25)	FUNCDATA	$1, gclocals·d20a545a455ba198dc85958c666abfdf(SB)
	0x002f 00047 (closure2.go:25)	FUNCDATA	$2, "".bar.stkobj(SB)
	0x002f 00047 (closure2.go:26)	PCDATA	$1, $0
	0x002f 00047 (closure2.go:26)	CALL	"".foo(SB)
	0x0034 00052 (closure2.go:26)	MOVQ	(SP), DX
	0x0038 00056 (closure2.go:26)	MOVQ	DX, "".f+88(SP)
	0x003d 00061 (closure2.go:27)	MOVQ	(DX), AX
	0x0040 00064 (closure2.go:27)	MOVQ	$5, (SP)
	0x0048 00072 (closure2.go:27)	PCDATA	$1, $1
	0x0048 00072 (closure2.go:27)	CALL	AX
	0x004a 00074 (closure2.go:28)	MOVQ	"".f+88(SP), AX
	0x004f 00079 (closure2.go:28)	MOVQ	AX, "".pc+80(SP)
	0x0054 00084 (closure2.go:29)	TESTB	AL, (AX)
	0x0056 00086 (closure2.go:29)	LEAQ	type."".closure(SB), CX
	0x005d 00093 (closure2.go:29)	MOVQ	CX, (SP)
	0x0061 00097 (closure2.go:29)	MOVQ	AX, 8(SP)
	0x0066 00102 (closure2.go:29)	PCDATA	$1, $2
	0x0066 00102 (closure2.go:29)	CALL	runtime.convT2E(SB)
	0x006b 00107 (closure2.go:29)	MOVQ	16(SP), AX
	0x0070 00112 (closure2.go:29)	MOVQ	24(SP), CX
	0x0075 00117 (closure2.go:29)	XORPS	X0, X0
	0x0078 00120 (closure2.go:29)	MOVUPS	X0, ""..autotmp_33+112(SP)
	0x007d 00125 (closure2.go:29)	MOVQ	AX, ""..autotmp_33+112(SP)
	0x0082 00130 (closure2.go:29)	MOVQ	CX, ""..autotmp_33+120(SP)
	0x0087 00135 (<unknown line number>)	NOP
	0x0087 00135 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), AX
	0x008e 00142 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), CX
	0x0095 00149 ($GOROOT/src/fmt/print.go:213)	MOVQ	CX, (SP)
	0x0099 00153 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 8(SP)
	0x009e 00158 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."%#v\n"(SB), AX
	0x00a5 00165 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 16(SP)
	0x00aa 00170 ($GOROOT/src/fmt/print.go:213)	MOVQ	$4, 24(SP)
	0x00b3 00179 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_33+112(SP), AX
	0x00b8 00184 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 32(SP)
	0x00bd 00189 ($GOROOT/src/fmt/print.go:213)	MOVQ	$1, 40(SP)
	0x00c6 00198 ($GOROOT/src/fmt/print.go:213)	MOVQ	$1, 48(SP)
	0x00cf 00207 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x00d4 00212 (closure2.go:30)	MOVQ	"".pc+80(SP), AX
	0x00d9 00217 (closure2.go:30)	MOVQ	8(AX), CX
	0x00dd 00221 (closure2.go:30)	MOVQ	(CX), CX
	0x00e0 00224 (closure2.go:30)	MOVQ	CX, (SP)
	0x00e4 00228 (closure2.go:30)	CALL	runtime.convT64(SB)
	0x00e9 00233 (closure2.go:30)	MOVQ	"".pc+80(SP), AX
	0x00ee 00238 (closure2.go:30)	MOVQ	16(AX), CX
	0x00f2 00242 (closure2.go:30)	MOVQ	8(SP), DX
	0x00f7 00247 (closure2.go:30)	MOVQ	DX, ""..autotmp_87+104(SP)
	0x00fc 00252 (closure2.go:30)	MOVQ	(CX), CX
	0x00ff 00255 (closure2.go:30)	MOVQ	CX, (SP)
	0x0103 00259 (closure2.go:30)	PCDATA	$1, $3
	0x0103 00259 (closure2.go:30)	CALL	runtime.convT64(SB)
	0x0108 00264 (closure2.go:30)	MOVQ	"".pc+80(SP), AX
	0x010d 00269 (closure2.go:30)	MOVQ	24(AX), CX
	0x0111 00273 (closure2.go:30)	MOVQ	8(SP), DX
	0x0116 00278 (closure2.go:30)	MOVQ	DX, ""..autotmp_88+96(SP)
	0x011b 00283 (closure2.go:30)	MOVQ	(CX), CX
	0x011e 00286 (closure2.go:30)	MOVQ	CX, (SP)
	0x0122 00290 (closure2.go:30)	PCDATA	$1, $4
	0x0122 00290 (closure2.go:30)	CALL	runtime.convT64(SB)
	0x0127 00295 (closure2.go:30)	MOVQ	8(SP), AX
	0x012c 00300 (closure2.go:30)	XORPS	X0, X0
	0x012f 00303 (closure2.go:30)	MOVUPS	X0, ""..autotmp_45+176(SP)
	0x0137 00311 (closure2.go:30)	MOVUPS	X0, ""..autotmp_45+192(SP)
	0x013f 00319 (closure2.go:30)	MOVUPS	X0, ""..autotmp_45+208(SP)
	0x0147 00327 (closure2.go:30)	LEAQ	type.int(SB), CX
	0x014e 00334 (closure2.go:30)	MOVQ	CX, ""..autotmp_45+176(SP)
	0x0156 00342 (closure2.go:30)	MOVQ	""..autotmp_87+104(SP), DX
	0x015b 00347 (closure2.go:30)	MOVQ	DX, ""..autotmp_45+184(SP)
	0x0163 00355 (closure2.go:30)	MOVQ	CX, ""..autotmp_45+192(SP)
	0x016b 00363 (closure2.go:30)	MOVQ	""..autotmp_88+96(SP), DX
	0x0170 00368 (closure2.go:30)	MOVQ	DX, ""..autotmp_45+200(SP)
	0x0178 00376 (closure2.go:30)	MOVQ	CX, ""..autotmp_45+208(SP)
	0x0180 00384 (closure2.go:30)	MOVQ	AX, ""..autotmp_45+216(SP)
	0x0188 00392 (<unknown line number>)	NOP
	0x0188 00392 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), AX
	0x018f 00399 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), DX
	0x0196 00406 ($GOROOT/src/fmt/print.go:213)	MOVQ	DX, (SP)
	0x019a 00410 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 8(SP)
	0x019f 00415 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a=%d, b=%d,c=%d\n"(SB), AX
	0x01a6 00422 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 16(SP)
	0x01ab 00427 ($GOROOT/src/fmt/print.go:213)	MOVQ	$16, 24(SP)
	0x01b4 00436 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_45+176(SP), BX
	0x01bc 00444 ($GOROOT/src/fmt/print.go:213)	MOVQ	BX, 32(SP)
	0x01c1 00449 ($GOROOT/src/fmt/print.go:213)	MOVQ	$3, 40(SP)
	0x01ca 00458 ($GOROOT/src/fmt/print.go:213)	MOVQ	$3, 48(SP)
	0x01d3 00467 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $2
	0x01d3 00467 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x01d8 00472 (closure2.go:31)	MOVQ	"".f+88(SP), DX
	0x01dd 00477 (closure2.go:31)	MOVQ	(DX), AX
	0x01e0 00480 (closure2.go:31)	MOVQ	$6, (SP)
	0x01e8 00488 (closure2.go:31)	PCDATA	$1, $5
	0x01e8 00488 (closure2.go:31)	CALL	AX
	0x01ea 00490 (closure2.go:32)	MOVQ	"".pc+80(SP), AX
	0x01ef 00495 (closure2.go:32)	MOVQ	8(AX), CX
	0x01f3 00499 (closure2.go:32)	MOVQ	(CX), CX
	0x01f6 00502 (closure2.go:32)	MOVQ	CX, (SP)
	0x01fa 00506 (closure2.go:32)	CALL	runtime.convT64(SB)
	0x01ff 00511 (closure2.go:32)	MOVQ	"".pc+80(SP), AX
	0x0204 00516 (closure2.go:32)	MOVQ	16(AX), CX
	0x0208 00520 (closure2.go:32)	MOVQ	8(SP), DX
	0x020d 00525 (closure2.go:32)	MOVQ	DX, ""..autotmp_87+104(SP)
	0x0212 00530 (closure2.go:32)	MOVQ	(CX), CX
	0x0215 00533 (closure2.go:32)	MOVQ	CX, (SP)
	0x0219 00537 (closure2.go:32)	PCDATA	$1, $6
	0x0219 00537 (closure2.go:32)	CALL	runtime.convT64(SB)
	0x021e 00542 (closure2.go:32)	MOVQ	"".pc+80(SP), AX
	0x0223 00547 (closure2.go:32)	MOVQ	24(AX), AX
	0x0227 00551 (closure2.go:32)	MOVQ	8(SP), CX
	0x022c 00556 (closure2.go:32)	MOVQ	CX, ""..autotmp_88+96(SP)
	0x0231 00561 (closure2.go:32)	MOVQ	(AX), AX
	0x0234 00564 (closure2.go:32)	MOVQ	AX, (SP)
	0x0238 00568 (closure2.go:32)	PCDATA	$1, $7
	0x0238 00568 (closure2.go:32)	CALL	runtime.convT64(SB)
	0x023d 00573 (closure2.go:32)	MOVQ	8(SP), AX
	0x0242 00578 (closure2.go:32)	XORPS	X0, X0
	0x0245 00581 (closure2.go:32)	MOVUPS	X0, ""..autotmp_57+128(SP)
	0x024d 00589 (closure2.go:32)	MOVUPS	X0, ""..autotmp_57+144(SP)
	0x0255 00597 (closure2.go:32)	MOVUPS	X0, ""..autotmp_57+160(SP)
	0x025d 00605 (closure2.go:32)	LEAQ	type.int(SB), CX
	0x0264 00612 (closure2.go:32)	MOVQ	CX, ""..autotmp_57+128(SP)
	0x026c 00620 (closure2.go:32)	MOVQ	""..autotmp_87+104(SP), DX
	0x0271 00625 (closure2.go:32)	MOVQ	DX, ""..autotmp_57+136(SP)
	0x0279 00633 (closure2.go:32)	MOVQ	CX, ""..autotmp_57+144(SP)
	0x0281 00641 (closure2.go:32)	MOVQ	""..autotmp_88+96(SP), DX
	0x0286 00646 (closure2.go:32)	MOVQ	DX, ""..autotmp_57+152(SP)
	0x028e 00654 (closure2.go:32)	MOVQ	CX, ""..autotmp_57+160(SP)
	0x0296 00662 (closure2.go:32)	MOVQ	AX, ""..autotmp_57+168(SP)
	0x029e 00670 (<unknown line number>)	NOP
	0x029e 00670 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), AX
	0x02a5 00677 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), CX
	0x02ac 00684 ($GOROOT/src/fmt/print.go:213)	MOVQ	CX, (SP)
	0x02b0 00688 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 8(SP)
	0x02b5 00693 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."a=%d, b=%d,c=%d\n"(SB), AX
	0x02bc 00700 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 16(SP)
	0x02c1 00705 ($GOROOT/src/fmt/print.go:213)	MOVQ	$16, 24(SP)
	0x02ca 00714 ($GOROOT/src/fmt/print.go:213)	LEAQ	""..autotmp_57+128(SP), AX
	0x02d2 00722 ($GOROOT/src/fmt/print.go:213)	MOVQ	AX, 32(SP)
	0x02d7 00727 ($GOROOT/src/fmt/print.go:213)	MOVQ	$3, 40(SP)
	0x02e0 00736 ($GOROOT/src/fmt/print.go:213)	MOVQ	$3, 48(SP)
	0x02e9 00745 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $0
	0x02e9 00745 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x02ee 00750 (closure2.go:32)	MOVQ	224(SP), BP
	0x02f6 00758 (closure2.go:32)	ADDQ	$232, SP
	0x02fd 00765 (closure2.go:32)	RET
	0x02fe 00766 (closure2.go:32)	NOP
	0x02fe 00766 (closure2.go:25)	PCDATA	$1, $-1
	0x02fe 00766 (closure2.go:25)	PCDATA	$0, $-2
	0x02fe 00766 (closure2.go:25)	NOP
	0x0300 00768 (closure2.go:25)	CALL	runtime.morestack_noctxt(SB)
	0x0305 00773 (closure2.go:25)	PCDATA	$0, $-1
	0x0305 00773 (closure2.go:25)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 98 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 e8 02 00 00 48 81 ec e8 00 00 00 48  A.......H......H
	0x0020 89 ac 24 e0 00 00 00 48 8d ac 24 e0 00 00 00 e8  ..$....H..$.....
	0x0030 00 00 00 00 48 8b 14 24 48 89 54 24 58 48 8b 02  ....H..$H.T$XH..
	0x0040 48 c7 04 24 05 00 00 00 ff d0 48 8b 44 24 58 48  H..$......H.D$XH
	0x0050 89 44 24 50 84 00 48 8d 0d 00 00 00 00 48 89 0c  .D$P..H......H..
	0x0060 24 48 89 44 24 08 e8 00 00 00 00 48 8b 44 24 10  $H.D$......H.D$.
	0x0070 48 8b 4c 24 18 0f 57 c0 0f 11 44 24 70 48 89 44  H.L$..W...D$pH.D
	0x0080 24 70 48 89 4c 24 78 48 8b 05 00 00 00 00 48 8d  $pH.L$xH......H.
	0x0090 0d 00 00 00 00 48 89 0c 24 48 89 44 24 08 48 8d  .....H..$H.D$.H.
	0x00a0 05 00 00 00 00 48 89 44 24 10 48 c7 44 24 18 04  .....H.D$.H.D$..
	0x00b0 00 00 00 48 8d 44 24 70 48 89 44 24 20 48 c7 44  ...H.D$pH.D$ H.D
	0x00c0 24 28 01 00 00 00 48 c7 44 24 30 01 00 00 00 e8  $(....H.D$0.....
	0x00d0 00 00 00 00 48 8b 44 24 50 48 8b 48 08 48 8b 09  ....H.D$PH.H.H..
	0x00e0 48 89 0c 24 e8 00 00 00 00 48 8b 44 24 50 48 8b  H..$.....H.D$PH.
	0x00f0 48 10 48 8b 54 24 08 48 89 54 24 68 48 8b 09 48  H.H.T$.H.T$hH..H
	0x0100 89 0c 24 e8 00 00 00 00 48 8b 44 24 50 48 8b 48  ..$.....H.D$PH.H
	0x0110 18 48 8b 54 24 08 48 89 54 24 60 48 8b 09 48 89  .H.T$.H.T$`H..H.
	0x0120 0c 24 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0 0f  .$.....H.D$..W..
	0x0130 11 84 24 b0 00 00 00 0f 11 84 24 c0 00 00 00 0f  ..$.......$.....
	0x0140 11 84 24 d0 00 00 00 48 8d 0d 00 00 00 00 48 89  ..$....H......H.
	0x0150 8c 24 b0 00 00 00 48 8b 54 24 68 48 89 94 24 b8  .$....H.T$hH..$.
	0x0160 00 00 00 48 89 8c 24 c0 00 00 00 48 8b 54 24 60  ...H..$....H.T$`
	0x0170 48 89 94 24 c8 00 00 00 48 89 8c 24 d0 00 00 00  H..$....H..$....
	0x0180 48 89 84 24 d8 00 00 00 48 8b 05 00 00 00 00 48  H..$....H......H
	0x0190 8d 15 00 00 00 00 48 89 14 24 48 89 44 24 08 48  ......H..$H.D$.H
	0x01a0 8d 05 00 00 00 00 48 89 44 24 10 48 c7 44 24 18  ......H.D$.H.D$.
	0x01b0 10 00 00 00 48 8d 9c 24 b0 00 00 00 48 89 5c 24  ....H..$....H.\$
	0x01c0 20 48 c7 44 24 28 03 00 00 00 48 c7 44 24 30 03   H.D$(....H.D$0.
	0x01d0 00 00 00 e8 00 00 00 00 48 8b 54 24 58 48 8b 02  ........H.T$XH..
	0x01e0 48 c7 04 24 06 00 00 00 ff d0 48 8b 44 24 50 48  H..$......H.D$PH
	0x01f0 8b 48 08 48 8b 09 48 89 0c 24 e8 00 00 00 00 48  .H.H..H..$.....H
	0x0200 8b 44 24 50 48 8b 48 10 48 8b 54 24 08 48 89 54  .D$PH.H.H.T$.H.T
	0x0210 24 68 48 8b 09 48 89 0c 24 e8 00 00 00 00 48 8b  $hH..H..$.....H.
	0x0220 44 24 50 48 8b 40 18 48 8b 4c 24 08 48 89 4c 24  D$PH.@.H.L$.H.L$
	0x0230 60 48 8b 00 48 89 04 24 e8 00 00 00 00 48 8b 44  `H..H..$.....H.D
	0x0240 24 08 0f 57 c0 0f 11 84 24 80 00 00 00 0f 11 84  $..W....$.......
	0x0250 24 90 00 00 00 0f 11 84 24 a0 00 00 00 48 8d 0d  $.......$....H..
	0x0260 00 00 00 00 48 89 8c 24 80 00 00 00 48 8b 54 24  ....H..$....H.T$
	0x0270 68 48 89 94 24 88 00 00 00 48 89 8c 24 90 00 00  hH..$....H..$...
	0x0280 00 48 8b 54 24 60 48 89 94 24 98 00 00 00 48 89  .H.T$`H..$....H.
	0x0290 8c 24 a0 00 00 00 48 89 84 24 a8 00 00 00 48 8b  .$....H..$....H.
	0x02a0 05 00 00 00 00 48 8d 0d 00 00 00 00 48 89 0c 24  .....H......H..$
	0x02b0 48 89 44 24 08 48 8d 05 00 00 00 00 48 89 44 24  H.D$.H......H.D$
	0x02c0 10 48 c7 44 24 18 10 00 00 00 48 8d 84 24 80 00  .H.D$.....H..$..
	0x02d0 00 00 48 89 44 24 20 48 c7 44 24 28 03 00 00 00  ..H.D$ H.D$(....
	0x02e0 48 c7 44 24 30 03 00 00 00 e8 00 00 00 00 48 8b  H.D$0.........H.
	0x02f0 ac 24 e0 00 00 00 48 81 c4 e8 00 00 00 c3 66 90  .$....H.......f.
	0x0300 e8 00 00 00 00 e9 f6 fc ff ff                    ..........
	rel 3+0 t=25 type.*os.File+0
	rel 3+0 t=25 type.*os.File+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.*os.File+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type."".closure+0
	rel 5+4 t=17 TLS+0
	rel 48+4 t=8 "".foo+0
	rel 72+0 t=11 +0
	rel 89+4 t=16 type."".closure+0
	rel 103+4 t=8 runtime.convT2E+0
	rel 138+4 t=16 os.Stdout+0
	rel 145+4 t=16 go.itab.*os.File,io.Writer+0
	rel 161+4 t=16 go.string."%#v\n"+0
	rel 208+4 t=8 fmt.Fprintf+0
	rel 229+4 t=8 runtime.convT64+0
	rel 260+4 t=8 runtime.convT64+0
	rel 291+4 t=8 runtime.convT64+0
	rel 330+4 t=16 type.int+0
	rel 395+4 t=16 os.Stdout+0
	rel 402+4 t=16 go.itab.*os.File,io.Writer+0
	rel 418+4 t=16 go.string."a=%d, b=%d,c=%d\n"+0
	rel 468+4 t=8 fmt.Fprintf+0
	rel 488+0 t=11 +0
	rel 507+4 t=8 runtime.convT64+0
	rel 538+4 t=8 runtime.convT64+0
	rel 569+4 t=8 runtime.convT64+0
	rel 608+4 t=16 type.int+0
	rel 673+4 t=16 os.Stdout+0
	rel 680+4 t=16 go.itab.*os.File,io.Writer+0
	rel 696+4 t=16 go.string."a=%d, b=%d,c=%d\n"+0
	rel 746+4 t=8 fmt.Fprintf+0
	rel 769+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=53 args=0x0 locals=0x8 funcid=0x0
	0x0000 00000 (closure2.go:35)	TEXT	"".main(SB), ABIInternal, $8-0
	0x0000 00000 (closure2.go:35)	MOVQ	(TLS), CX
	0x0009 00009 (closure2.go:35)	CMPQ	SP, 16(CX)
	0x000d 00013 (closure2.go:35)	PCDATA	$0, $-2
	0x000d 00013 (closure2.go:35)	JLS	46
	0x000f 00015 (closure2.go:35)	PCDATA	$0, $-1
	0x000f 00015 (closure2.go:35)	SUBQ	$8, SP
	0x0013 00019 (closure2.go:35)	MOVQ	BP, (SP)
	0x0017 00023 (closure2.go:35)	LEAQ	(SP), BP
	0x001b 00027 (closure2.go:35)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (closure2.go:35)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (closure2.go:36)	PCDATA	$1, $0
	0x001b 00027 (closure2.go:36)	NOP
	0x0020 00032 (closure2.go:36)	CALL	"".bar(SB)
	0x0025 00037 (closure2.go:37)	MOVQ	(SP), BP
	0x0029 00041 (closure2.go:37)	ADDQ	$8, SP
	0x002d 00045 (closure2.go:37)	RET
	0x002e 00046 (closure2.go:37)	NOP
	0x002e 00046 (closure2.go:35)	PCDATA	$1, $-1
	0x002e 00046 (closure2.go:35)	PCDATA	$0, $-2
	0x002e 00046 (closure2.go:35)	CALL	runtime.morestack_noctxt(SB)
	0x0033 00051 (closure2.go:35)	PCDATA	$0, $-1
	0x0033 00051 (closure2.go:35)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 1f 48  eH..%....H;a.v.H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f 1f 44 00 00  ...H.,$H.,$..D..
	0x0020 e8 00 00 00 00 48 8b 2c 24 48 83 c4 08 c3 e8 00  .....H.,$H......
	0x0030 00 00 00 eb cb                                   .....
	rel 5+4 t=17 TLS+0
	rel 33+4 t=8 "".bar+0
	rel 47+4 t=8 runtime.morestack_noctxt+0
"".foo.func1 STEXT nosplit size=48 args=0x10 locals=0x0 funcid=0x0
	0x0000 00000 (closure2.go:10)	TEXT	"".foo.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $0-16
	0x0000 00000 (closure2.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure2.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure2.go:10)	MOVQ	16(DX), AX
	0x0004 00004 (closure2.go:10)	MOVQ	24(DX), CX
	0x0008 00008 (closure2.go:10)	MOVQ	8(DX), DX
	0x000c 00012 (closure2.go:11)	MOVQ	"".n+8(SP), BX
	0x0011 00017 (closure2.go:11)	ADDQ	BX, (DX)
	0x0014 00020 (closure2.go:12)	ADDQ	BX, (AX)
	0x0017 00023 (closure2.go:13)	MOVQ	(CX), SI
	0x001a 00026 (closure2.go:13)	ADDQ	SI, BX
	0x001d 00029 (closure2.go:13)	MOVQ	BX, (CX)
	0x0020 00032 (closure2.go:14)	MOVQ	(DX), CX
	0x0023 00035 (closure2.go:14)	ADDQ	(AX), CX
	0x0026 00038 (closure2.go:14)	LEAQ	(CX)(BX*1), AX
	0x002a 00042 (closure2.go:14)	MOVQ	AX, "".~r1+16(SP)
	0x002f 00047 (closure2.go:14)	RET
	0x0000 48 8b 42 10 48 8b 4a 18 48 8b 52 08 48 8b 5c 24  H.B.H.J.H.R.H.\$
	0x0010 08 48 01 1a 48 01 18 48 8b 31 48 01 f3 48 89 19  .H..H..H.1H..H..
	0x0020 48 8b 0a 48 03 08 48 8d 04 19 48 89 44 24 10 c3  H..H..H...H.D$..
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
type..eq.[3]interface {} STEXT dupok size=170 args=0x18 locals=0x30 funcid=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq.[3]interface {}(SB), DUPOK|ABIInternal, $48-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	160
	0x0013 00019 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0013 00019 (<autogenerated>:1)	SUBQ	$48, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 40(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	40(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (<autogenerated>:1)	MOVQ	"".q+64(SP), AX
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+56(SP), CX
	0x002b 00043 (<autogenerated>:1)	XORL	DX, DX
	0x002d 00045 (<autogenerated>:1)	JMP	66
	0x002f 00047 (<autogenerated>:1)	MOVQ	""..autotmp_6+32(SP), BX
	0x0034 00052 (<autogenerated>:1)	LEAQ	1(BX), DX
	0x0038 00056 (<autogenerated>:1)	MOVQ	"".q+64(SP), AX
	0x003d 00061 (<autogenerated>:1)	MOVQ	"".p+56(SP), CX
	0x0042 00066 (<autogenerated>:1)	CMPQ	DX, $3
	0x0046 00070 (<autogenerated>:1)	JGE	149
	0x0048 00072 (<autogenerated>:1)	MOVQ	DX, BX
	0x004b 00075 (<autogenerated>:1)	SHLQ	$4, DX
	0x004f 00079 (<autogenerated>:1)	MOVQ	(CX)(DX*1), SI
	0x0053 00083 (<autogenerated>:1)	MOVQ	(AX)(DX*1), DI
	0x0057 00087 (<autogenerated>:1)	MOVQ	8(DX)(CX*1), R8
	0x005c 00092 (<autogenerated>:1)	MOVQ	8(DX)(AX*1), DX
	0x0061 00097 (<autogenerated>:1)	CMPQ	DI, SI
	0x0064 00100 (<autogenerated>:1)	JNE	133
	0x0066 00102 (<autogenerated>:1)	MOVQ	BX, ""..autotmp_6+32(SP)
	0x006b 00107 (<autogenerated>:1)	MOVQ	SI, (SP)
	0x006f 00111 (<autogenerated>:1)	MOVQ	R8, 8(SP)
	0x0074 00116 (<autogenerated>:1)	MOVQ	DX, 16(SP)
	0x0079 00121 (<autogenerated>:1)	PCDATA	$1, $0
	0x0079 00121 (<autogenerated>:1)	CALL	runtime.efaceeq(SB)
	0x007e 00126 (<autogenerated>:1)	CMPB	24(SP), $0
	0x0083 00131 (<autogenerated>:1)	JNE	47
	0x0085 00133 (<autogenerated>:1)	XORL	AX, AX
	0x0087 00135 (<autogenerated>:1)	MOVB	AL, "".r+72(SP)
	0x008b 00139 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x0090 00144 (<autogenerated>:1)	ADDQ	$48, SP
	0x0094 00148 (<autogenerated>:1)	RET
	0x0095 00149 (<autogenerated>:1)	MOVL	$1, AX
	0x009a 00154 (<autogenerated>:1)	JMP	135
	0x009c 00156 (<autogenerated>:1)	NOP
	0x009c 00156 (<autogenerated>:1)	PCDATA	$1, $-1
	0x009c 00156 (<autogenerated>:1)	PCDATA	$0, $-2
	0x009c 00156 (<autogenerated>:1)	NOP
	0x00a0 00160 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00a5 00165 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00a5 00165 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 8d  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
	0x0020 28 48 8b 44 24 40 48 8b 4c 24 38 31 d2 eb 13 48  (H.D$@H.L$81...H
	0x0030 8b 5c 24 20 48 8d 53 01 48 8b 44 24 40 48 8b 4c  .\$ H.S.H.D$@H.L
	0x0040 24 38 48 83 fa 03 7d 4d 48 89 d3 48 c1 e2 04 48  $8H...}MH..H...H
	0x0050 8b 34 11 48 8b 3c 10 4c 8b 44 0a 08 48 8b 54 02  .4.H.<.L.D..H.T.
	0x0060 08 48 39 f7 75 1f 48 89 5c 24 20 48 89 34 24 4c  .H9.u.H.\$ H.4$L
	0x0070 89 44 24 08 48 89 54 24 10 e8 00 00 00 00 80 7c  .D$.H.T$.......|
	0x0080 24 18 00 75 aa 31 c0 88 44 24 48 48 8b 6c 24 28  $..u.1..D$HH.l$(
	0x0090 48 83 c4 30 c3 b8 01 00 00 00 eb eb 0f 1f 40 00  H..0..........@.
	0x00a0 e8 00 00 00 00 e9 56 ff ff ff                    ......V...
	rel 5+4 t=17 TLS+0
	rel 122+4 t=8 runtime.efaceeq+0
	rel 161+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.info.fmt.Printf$abstract SDWARFABSFCN dupok size=54
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 66 00 01 01 11 66  .fmt.Printf....f
	0x0010 6f 72 6d 61 74 00 00 00 00 00 00 11 61 00 00 00  ormat.......a...
	0x0020 00 00 00 11 6e 00 01 00 00 00 00 11 65 72 72 00  ....n.......err.
	0x0030 01 00 00 00 00 00                                ......
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.error+0
	rel 0+0 t=24 type.int+0
	rel 0+0 t=24 type.string+0
	rel 23+4 t=31 go.info.string+0
	rel 31+4 t=31 go.info.[]interface {}+0
	rel 39+4 t=31 go.info.int+0
	rel 49+4 t=31 go.info.error+0
go.string."%#v\n" SRODATA dupok size=4
	0x0000 25 23 76 0a                                      %#v.
go.string."a=%d, b=%d,c=%d\n" SRODATA dupok size=16
	0x0000 61 3d 25 64 2c 20 62 3d 25 64 2c 63 3d 25 64 0a  a=%d, b=%d,c=%d.
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*func(int) int- SRODATA dupok size=17
	0x0000 00 00 0e 2a 66 75 6e 63 28 69 6e 74 29 20 69 6e  ...*func(int) in
	0x0010 74                                               t
type.*func(int) int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2f eb d1 1f 08 08 08 36 00 00 00 00 00 00 00 00  /......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int) int-+0
	rel 48+8 t=1 type.func(int) int+0
type.func(int) int SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 98 3c 32 87 02 08 08 33 00 00 00 00 00 00 00 00  .<2....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int) int-+0
	rel 44+4 t=6 type.*func(int) int+0
	rel 56+8 t=1 type.int+0
	rel 64+8 t=1 type.int+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
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
type..eqfunc.[3]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[3]interface {}+0
type..namedata.*[3]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 33 5d 69 6e 74 65 72 66 61 63 65  ...*[3]interface
	0x0010 20 7b 7d                                          {}
type.*[3]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b0 18 fe b9 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[3]interface {}-+0
	rel 48+8 t=1 type.[3]interface {}+0
runtime.gcbits.2a SRODATA dupok size=1
	0x0000 2a                                               *
type.[3]interface {} SRODATA dupok size=72
	0x0000 30 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  0.......0.......
	0x0010 1d dd cf d9 02 08 08 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc.[3]interface {}+0
	rel 32+8 t=1 runtime.gcbits.2a+0
	rel 40+4 t=5 type..namedata.*[3]interface {}-+0
	rel 44+4 t=6 type.*[3]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..eqfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.**main.closure- SRODATA dupok size=17
	0x0000 00 00 0e 2a 2a 6d 61 69 6e 2e 63 6c 6f 73 75 72  ...**main.closur
	0x0010 65                                               e
type.**"".closure SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c4 93 90 ca 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**main.closure-+0
	rel 48+8 t=1 type.*"".closure+0
type..namedata.*main.closure- SRODATA dupok size=16
	0x0000 00 00 0d 2a 6d 61 69 6e 2e 63 6c 6f 73 75 72 65  ...*main.closure
type.*"".closure SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 3e c8 f6 43 08 08 08 36 00 00 00 00 00 00 00 00  >..C...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.closure-+0
	rel 44+4 t=6 type.**"".closure+0
	rel 48+8 t=1 type."".closure+0
runtime.gcbits.0e SRODATA dupok size=1
	0x0000 0e                                               .
type..namedata.f- SRODATA dupok size=4
	0x0000 00 00 01 66                                      ...f
type..namedata.a- SRODATA dupok size=4
	0x0000 00 00 01 61                                      ...a
type..namedata.b- SRODATA dupok size=4
	0x0000 00 00 01 62                                      ...b
type..namedata.c- SRODATA dupok size=4
	0x0000 00 00 01 63                                      ...c
type."".closure SRODATA size=192
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 bb 0b bc 3f 0f 08 08 19 00 00 00 00 00 00 00 00  ...?............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 70 00 00 00 00 00 00 00  ........p.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x00b0 00 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  ........0.......
	rel 24+8 t=1 type..eqfunc32+0
	rel 32+8 t=1 runtime.gcbits.0e+0
	rel 40+4 t=5 type..namedata.*main.closure-+0
	rel 44+4 t=5 type.*"".closure+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".closure+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.f-+0
	rel 104+8 t=1 type.uintptr+0
	rel 120+8 t=1 type..namedata.a-+0
	rel 128+8 t=1 type.*int+0
	rel 144+8 t=1 type..namedata.b-+0
	rel 152+8 t=1 type.*int+0
	rel 168+8 t=1 type..namedata.c-+0
	rel 176+8 t=1 type.*int+0
type..namedata.*struct { F uintptr; a *int; b *int; c *int }- SRODATA dupok size=48
	0x0000 00 00 2d 2a 73 74 72 75 63 74 20 7b 20 46 20 75  ..-*struct { F u
	0x0010 69 6e 74 70 74 72 3b 20 61 20 2a 69 6e 74 3b 20  intptr; a *int; 
	0x0020 62 20 2a 69 6e 74 3b 20 63 20 2a 69 6e 74 20 7d  b *int; c *int }
type.*struct { F uintptr; "".a *int; "".b *int; "".c *int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 46 4d 41 18 08 08 08 36 00 00 00 00 00 00 00 00  FMA....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; a *int; b *int; c *int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; "".a *int; "".b *int; "".c *int }+0
type..namedata..F- SRODATA dupok size=5
	0x0000 00 00 02 2e 46                                   ....F
type.noalg.struct { F uintptr; "".a *int; "".b *int; "".c *int } SRODATA dupok size=176
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 c9 6d 98 d2 02 08 08 19 00 00 00 00 00 00 00 00  .m..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x00a0 00 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  ........0.......
	rel 32+8 t=1 runtime.gcbits.0e+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; a *int; b *int; c *int }-+0
	rel 44+4 t=6 type.*struct { F uintptr; "".a *int; "".b *int; "".c *int }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; "".a *int; "".b *int; "".c *int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.a-+0
	rel 112+8 t=1 type.*int+0
	rel 128+8 t=1 type..namedata.b-+0
	rel 136+8 t=1 type.*int+0
	rel 152+8 t=1 type..namedata.c-+0
	rel 160+8 t=1 type.*int+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=1 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.unsafe. SRODATA dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·5f7ae22b544db82d5d4c812af83655e9 SRODATA dupok size=12
	0x0000 04 00 00 00 01 00 00 00 00 00 00 00              ............
gclocals·c8a67faffb154eea2f3f8b30fe59d7e4 SRODATA dupok size=12
	0x0000 04 00 00 00 03 00 00 00 00 04 06 07              ............
gclocals·0ce64bbc7cfa5ef04d41c861de81a3d7 SRODATA dupok size=8
	0x0000 08 00 00 00 00 00 00 00                          ........
gclocals·d20a545a455ba198dc85958c666abfdf SRODATA dupok size=32
	0x0000 08 00 00 00 12 00 00 00 00 00 00 02 00 00 03 00  ................
	0x0010 00 0b 00 00 0f 00 00 01 00 00 09 00 00 0c 00 00  ................
"".bar.stkobj SRODATA static size=72
	0x0000 04 00 00 00 00 00 00 00 78 ff ff ff ff ff ff ff  ........x.......
	0x0010 00 00 00 00 00 00 00 00 90 ff ff ff ff ff ff ff  ................
	0x0020 00 00 00 00 00 00 00 00 a0 ff ff ff ff ff ff ff  ................
	0x0030 00 00 00 00 00 00 00 00 d0 ff ff ff ff ff ff ff  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.func(int) int+0
	rel 32+8 t=1 type.[1]interface {}+0
	rel 48+8 t=1 type.[3]interface {}+0
	rel 64+8 t=1 type.[3]interface {}+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........

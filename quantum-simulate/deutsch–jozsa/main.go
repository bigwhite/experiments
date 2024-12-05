package main

import (
	"fmt"

	"github.com/itsubaki/q"
)

// 实现常数函数的oracle
func constantOracle(qsim *q.Q, controls []q.Qubit, target q.Qubit) {
	// 对于常数函数，这里什么都不做
	// 因为常数函数要么总是返回0，要么总是返回1
}

// 实现平衡函数的oracle
func balancedOracle(qsim *q.Q, controls []q.Qubit, target q.Qubit) {
	// 对于平衡函数，我们对输入做CNOT操作
	for _, control := range controls {
		qsim.CNOT(control, target)
	}
}

func deutschJozsa(n int, isConstant bool) string {
	// 创建量子模拟器
	qsim := q.New()

	// 创建n+1个量子比特的寄存器
	// n个输入比特加1个输出比特
	qreg := make([]q.Qubit, n+1)
	for i := range qreg {
		qreg[i] = qsim.Zero()
	}

	// 将输出比特置为|1>
	qsim.X(qreg[n])

	// 对所有量子比特应用H门
	for _, qubit := range qreg {
		qsim.H(qubit)
	}

	// 应用oracle
	if isConstant {
		constantOracle(qsim, qreg[:n], qreg[n])
	} else {
		balancedOracle(qsim, qreg[:n], qreg[n])
	}

	// 对输入寄存器应用H门
	for i := 0; i < n; i++ {
		qsim.H(qreg[i])
	}

	// 测量输入寄存器
	result := ""
	for i := 0; i < n; i++ {
		m := qsim.Measure(qreg[i])
		result += m.BinaryString()
	}

	return result
}

func main() {
	// 测试5个输入比特的情况
	n := 5

	// 测试常数函数
	resultConstant := deutschJozsa(n, true)
	fmt.Printf("常数函数的结果: %v\n", resultConstant)
	if resultConstant == "00000" {
		fmt.Println("检测到常数函数!")
	}

	// 测试平衡函数
	resultBalanced := deutschJozsa(n, false)
	fmt.Printf("平衡函数的结果: %v\n", resultBalanced)
	if resultBalanced != "00000" {
		fmt.Println("检测到平衡函数!")
	}
}

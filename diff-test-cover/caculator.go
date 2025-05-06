package calculator

import "fmt"

// Calculate 执行简单的算术运算
func Calculate(op string, a, b int) (int, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		// !!! Bug introduced here: should be a * b !!!
		fmt.Println("Executing multiplication logic...") // 添加打印以便观察
		return a + b, nil                                // 错误地执行了加法
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

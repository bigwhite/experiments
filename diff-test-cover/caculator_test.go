package calculator

import "testing"

func TestCalculateAdd(t *testing.T) {
	result, err := Calculate("add", 5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("add(5, 3) = %d; want 8", result)
	}
}

func TestCalculateSub(t *testing.T) {
	result, err := Calculate("sub", 5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("sub(5, 3) = %d; want 2", result)
	}
}

// 这个测试会因为 Bug 而失败
func TestCalculateMul(t *testing.T) {
	result, err := Calculate("mul", 5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 期望 15，但因为 Bug 实际返回 8
	if result != 15 {
		t.Errorf("mul(5, 3) = %d; want 15", result)
	}
}

package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Add(a, b int) int {
	return a + b
}

func TestRequire(t *testing.T) {
	// Equal断言
	require.Equal(t, 4, Add(1, 3), "The result should be 4")

	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	sl3 := []int{2, 3, 4}
	require.Equal(t, sl1, sl2, "sl1 should equal to sl2 ")

	p1 := &sl1
	p2 := &sl2
	require.Equal(t, p1, p2, "the content which p1 point to should equal to which p2 point to")

	err := errors.New("demo error")
	require.EqualError(t, err, "demo error")

	// require.Exactly(t, int32(123), int64(123)) // failed! both type and value must be same

	// 布尔断言
	require.True(t, 1+1 == 2, "1+1 == 2 should be true")
	require.Contains(t, "Hello World", "World")
	require.Contains(t, []string{"Hello", "World"}, "World")
	require.Contains(t, map[string]string{"Hello": "World"}, "Hello")
	require.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})

	// 反向断言
	require.NotEqual(t, 4, Add(2, 3), "The result should not be 4")
	require.NotEqual(t, sl1, sl3, "sl1 should not equal to sl3 ")
	require.False(t, 1+1 == 3, "1+1 == 3 should be false")
	require.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond) //1秒之内condition参数都不为true，每10毫秒检查一次
	require.NotContains(t, "Hello World", "Go")
}

func TestAdd(t *testing.T) {
	result := Add(1, 3)
	require.Equal(t, 4, result, "The result should be 4")
}

func TestAdd1(t *testing.T) {
	result := Add(1, 3)
	require.Equal(t, 4, result, "The result should be 4")
	result = Add(2, 2)
	require.Equal(t, 4, result, "The result should be 4")
	result = Add(2, 3)
	require.Equal(t, 5, result, "The result should be 5")
	result = Add(0, 3)
	require.Equal(t, 3, result, "The result should be 3")
	result = Add(-1, 1)
	require.Equal(t, 0, result, "The result should be 0")
}

func TestAdd2(t *testing.T) {
	require := require.New(t)

	result := Add(1, 3)
	require.Equal(4, result, "The result should be 4")
	result = Add(2, 2)
	require.Equal(4, result, "The result should be 4")
	result = Add(2, 3)
	require.Equal(5, result, "The result should be 5")
	result = Add(0, 3)
	require.Equal(3, result, "The result should be 3")
	result = Add(-1, 1)
	require.Equal(0, result, "The result should be 0")
}

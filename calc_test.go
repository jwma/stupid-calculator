package stupid_calculator

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 1

	// 期望结果
	expected := a + b

	// 使用 Add() 计算结果
	result := Add(a, b)

	// 如果 Add() 计算结果不等于我们的期望结果，则测试不通过
	if expected != result {
		t.Errorf("expected %d but got %d\n", expected, result)
	}
}

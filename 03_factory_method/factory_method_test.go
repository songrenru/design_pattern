package factory_method

import (
	"testing"
)

func compute(f OperatorFactory, a, b int) int {
	operator := f.Create()
	operator.SetA(a)
	operator.SetB(b)
	return operator.Result()
}

func TestOperator(t *testing.T) {
	var f OperatorFactory
	f = PlusOperatorFactory{}
	get := compute(f, 1, 2)
	want := 3
	if get != want {
		t.Fatal("error with factory method pattern")
	}

	f = MinusOperatorFactory{}
	get = compute(f, 2, 99)
	want = -97
	if get != want {
		t.Fatal("error with factory method pattern")
	}
}

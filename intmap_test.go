package utils

import (
	"testing"
)

type intMapTestCase struct {
	data     []int
	answer   []int
	function func(x int) int
}

func TestIntMap(t *testing.T) {
	f1 := func(x int) int { return x + 1 }
	f2 := func(x int) int { return x*x + 2*x + 1 }

	testcases := []intMapTestCase{
		intMapTestCase{[]int{1, 2, 3}, []int{2, 3, 4}, f1},
		intMapTestCase{[]int{1, 2, 3}, []int{4, 9, 16}, f2},
	}

	for _, v := range testcases {
		r := IntMap(v.function, v.data)
		expected := v.answer

		if !EqIntSlices(r, expected) {
			t.Errorf("TestIntMap failed: expected %v, result %v", expected, r)
		}
	}
}

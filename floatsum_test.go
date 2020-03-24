package utils

import "testing"

type floatSumTestCase struct {
	data   []float64
	answer float64
}

func TestFloatSum(t *testing.T) {
	testcases := []floatSumTestCase{
		floatSumTestCase{[]float64{1.3, 2.2}, 3.5},
		floatSumTestCase{[]float64{0.0, 0.0, 0.0, 0.0}, 0.0},
		floatSumTestCase{[]float64{-1.5, 2.5}, 1.0},
		floatSumTestCase{[]float64{-4.7, -3.3}, -8.0},
		floatSumTestCase{[]float64{1, 2, 3, -1, 2}, 7},
	}

	for _, v := range testcases {
		r := FloatSum(v.data...)
		expected := v.answer

		if r != expected {
			t.Errorf("TestFloatSum failed: expected %v, result %v", expected, r)
		}
	}
}

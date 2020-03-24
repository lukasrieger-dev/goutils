package utils

import (
	"testing"
)

type intSumTestCase struct {
	title  string
	data   []int
	answer int
}

// This is a table test
func TestIntSum(t *testing.T) {
	testcases := []intSumTestCase{
		intSumTestCase{"Test 1", []int{1, 2}, 3},
		intSumTestCase{"Test 2", []int{0, 0}, 0},
		intSumTestCase{"Test 3", []int{-1, 2}, 1},
		intSumTestCase{"Test 4", []int{-4, -3}, -7},
		intSumTestCase{"Test 5", []int{1, 2, 3, -1, 2}, 7},
	}

	for _, test := range testcases {
		t.Run(test.title, func(t *testing.T) {
			testSum(t, test)
		})
	}
}

func testSum(t *testing.T, test intSumTestCase) {
	title := test.title
	data := test.data
	expected := test.answer

	if got := IntSum(data...); expected != got {
		t.Errorf("Test %v failed. Expected %v, but got %v", title, expected, got)
	}
}

// A benchmark test. It can be launched with go test -bench .
func BenchmarkIntSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSum(i, i, i, i, i, i, i, i)
	}
}

// go test -v to see all test results
// go test -cover to get code coverage
// go test -coverprofile c.out writes results to file c.out
// go tool cover -html=c.out visualizes the results from c.out
// godoc -http=:8080

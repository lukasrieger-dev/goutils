package utils

import "testing"

type potTest struct {
	title string
	exp	int
	answer int
}

func testPowerOfTwo(t *testing.T) {
	tests := []potTest {
		potTest{
			title: "2 ^ 0",
			exp: 0,
			answer: 1,
		},
		potTest {
			title: "2 ^ 1",
			exp: 1,
			answer: 2,
		},
		potTest {
			title: "2 ^ 2",
			exp: 2,
			answer: 4,
		},
		potTest {
			title: "2 ^ 3",
			exp: 3,
			answer: 8,
		},
	}

	for _, v := range tests {
		title := v.title
		exponent := v.exp
		expected := v.answer
	
		if got := PowerOfTwo(exponent); expected != got {
			t.Errorf("Test %v failed. Expected %v, but got %v", title, expected, got)
		}
	}

}
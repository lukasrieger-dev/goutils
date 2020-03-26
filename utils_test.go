// Utility functions for go.
// They can be used for int, float64 and string data types.
// Author: Lukas Rieger
// Change Log
// 2020-03-24	Creation

package utils

import (
	"reflect"
	"testing"
)

func TestPowerString_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		p    PowerString
		want bool
	}{
		{"palindrome", "xouuox", true},
		{"not_palindrome", "apple", false},
		{"empty_string", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.isPalindrome(); got != tt.want {
				t.Errorf("PowerString.isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowerString_reverse(t *testing.T) {
	tests := []struct {
		name string
		p    PowerString
		want string
	}{
		{"reverse", "cow", "woc"},
		{"empty_string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.reverse()
			if string(tt.p) != tt.want {
				t.Errorf("PowerString.isPalindrome() = %v, want %v", tt.p, tt.want)
			}
		})
	}
}

func TestPowerOfTwo(t *testing.T) {
	tests := []struct {
		title  string
		exp    int
		answer int
	}{
		{"2 ^ -1", -1, 0},
		{"2 ^ 0", 0, 1},
		{"2 ^ 1", 1, 2},
		{"2 ^ 2", 2, 4},
		{"2 ^ 3", 3, 8},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			title := test.title
			exponent := test.exp
			wanted := test.answer

			if got := PowerOfTwo(exponent); wanted != got {
				t.Errorf("Test %v failed. Expected %v, but got %v", title, wanted, got)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		f  func(x float64) float64
		xi []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"map_inc", args{func(x float64) float64 { return x + 1.0 }, []float64{1, 2, 3}}, []float64{2.0, 3.0, 4.0}},
		{"map_dec", args{func(x float64) float64 { return x - 1.0 }, []float64{1, 2, 3}}, []float64{0.0, 1.0, 2.0}},
		{"map_sqr", args{func(x float64) float64 { return x * x }, []float64{1, 2, 3}}, []float64{1.0, 4.0, 9.0}},
		{"map_empty", args{func(x float64) float64 { return x + 1.0 }, []float64{}}, []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.f, tt.args.xi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		f  func(x float64) bool
		xi []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"filter_gt_one", args{func(x float64) bool { return x > 1 }, []float64{1, 2, 3}}, []float64{2.0, 3.0}},
		{"filter_ident", args{func(x float64) bool { return x == x }, []float64{1, 2, 3}}, []float64{1.0, 2.0, 3.0}},
		{"filter_empty", args{func(x float64) bool { return x > 3 }, []float64{}}, []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.f, tt.args.xi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		xi []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"sum", args{[]float64{3.0, -1.0, 5.0}}, 7.0},
		{"empty", args{[]float64{}}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.xi...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"gcd1", args{20, 5}, 5},
		{"gcd2", args{21, 28}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreatestCommonDivisor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"lcm1", args{20, 5}, 20},
		{"lcm2", args{3, -7}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestCommonMultiple(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

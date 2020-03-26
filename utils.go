// Utility functions for go.
// They can be used for int, float64 and string data types.
// Author: Lukas Rieger
// Change Log
// 2020-03-24	Creation

package utils

import (
	"strings"
)

// String with utility methods
type PowerString string

func (p *PowerString) reverse() {
	runes := []rune(*p)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*p = PowerString(string(runes))
}

// Check if this string is a palindrome.
func (p PowerString) isPalindrome() bool {
	original := string(p)
	p.reverse()
	reversed := string(p)
	return original == reversed
}

// Return the index of the substring in this string.
func (p PowerString) index(substring PowerString) int {
	return strings.Index(string(p), string(substring))
}

// Check if this string contains the given substring
func (p PowerString) contains(substring PowerString) bool {
	return strings.Contains(string(p), string(substring))
}

// Sums up all given values (float64).
func Sum(xi ...float64) float64 {
	total := 0.0

	for _, v := range xi {
		total += v
	}
	return total
}

// Returns a slice of integers containing only the
// int values from the input slice for which the function
// f yielded true.
func Filter(f func(x float64) bool, xi []float64) []float64 {
	if xi == nil || len(xi) == 0 {
		return xi
	}

	var results []float64
	for _, e := range xi {
		if f(e) {
			results = append(results, e)
		}
	}
	return results
}

// Maps a function to each element in a slice of int
func Map(f func(x float64) float64, xi []float64) []float64 {
	if xi == nil || len(xi) == 0 {
		return xi
	}

	var results []float64
	for _, e := range xi {
		results = append(results, f(e))
	}
	return results
}

// Returns 2 to the power of the given exponent.
func PowerOfTwo(exponent int) int {
	if exponent < 0 {
		return 0
	}
	return 1 << exponent
}

// Computes the greatest common divisor of a and b
func GreatestCommonDivisor(a, b int) int {
	var gcdnum int
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcdnum = i
		}
	}
	if gcdnum == 0 {
		gcdnum = 1
	}
	return gcdnum
}

// Computes the lowest common multiple of a and b
func LowestCommonMultiple(a, b int) int {
	var lcmnum int = 1

	if a == 0 || b == 0 {
		return 0
	}

	if a > b {
		lcmnum = a
	} else {
		lcmnum = b
	}
	for {
		if lcmnum%a == 0 && lcmnum%b == 0 {
			break
		}
		lcmnum++
	}
	return lcmnum
}

// go test -v to see all test results
// go test -cover to get code coverage
// go test -coverprofile c.out writes results to file c.out
// go tool cover -html=c.out visualizes the results from c.out
// godoc -http=:8080

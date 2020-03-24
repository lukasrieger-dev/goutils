// Utility functions for go.
// They can be used for int, float64 and string data types.
// Author: Lukas Rieger
// Change Log
// 2020-03-24	Creation

package utils

// Sums up all given values (float64).
func FloatSum(xi ...float64) float64 {
	total := 0.0

	for _, v := range xi {
		total += v
	}
	return total
}

// Sums up all given values (int).
func IntSum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

// Returns a slice of integers containing only the 
// int values from the input slice for which the function
// f yielded true.
func IntFilter(f func(x int) bool, xi []int) []int {
	if xi == nil || len(xi) == 0 {
		return xi
	}

	var results []int
	for _, e := range xi {
		if f(e) {
			results = append(results, e)
		}
	}
	return results
}


// Maps a function to each element in a slice of int
func IntMap(f func(x int) int, xi []int) []int {
	if xi == nil || len(xi) == 0 {
		return xi
	}

	var results []int
	for _, e := range xi {
		results = append(results, f(e))
	}
	return results
}

// Compare if two slices are equal
func EqIntSlices(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Returns 2 to the power of the given exponent.
func PowerOfTwo(exponent int) int {
	return  1 << exponent
}

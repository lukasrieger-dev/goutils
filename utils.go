package utils

/*
	Sums up all values in a slice
*/
func FloatSum(xi ...float64) float64 {
	total := 0.0

	for _, v := range xi {
		total += v
	}
	return total
}

/*
	Sums up all values in a slice
*/
func IntSum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

/*
	Filters a slice of int values
*/
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

/*
	Maps a function to each element in a slice of int
*/
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

/*
	Compare if two slices are equal
*/
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

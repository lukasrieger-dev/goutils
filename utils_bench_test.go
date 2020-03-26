package utils

import (
	"testing"
)

func benchmarkSum(b *testing.B, sum func(xi ...float64) float64, values []float64) {
	for i := 0; i < b.N; i++ {
		sum(values...)
	}
}

func BenchmarkSum1(b *testing.B) { benchmarkSum(b, Sum, []float64{1.0, 2.0, 3.0, 4.0, 5.0}) }
func BenchmarkSum2(b *testing.B) { benchmarkSum(b, Sum, []float64{-1.0, -2.0, -3.0, -4.0, -5.0}) }
func BenchmarkSum3(b *testing.B) { benchmarkSum(b, Sum, []float64{0.0, 0.0}) }
func BenchmarkSum4(b *testing.B) { benchmarkSum(b, Sum, []float64{9999999.0001, 9999.99}) }

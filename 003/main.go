package main

import (
	"fmt"
	"math"
)

func isPrime(n float64) bool {
	var i float64
	for i = 2; i <= math.Sqrt(n); i++ {
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}

func getFactors(n float64) []float64 {
	factors := make([]float64, 0)
	var i float64
	for i = math.Round(math.Sqrt(n)); i >= 2; i-- {
		if math.Mod(n, i) == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func LargestPrimeFactor(n float64) float64 {
	factors := getFactors(n)
	fmt.Println(factors)
	for _, f := range factors {
		if isPrime(f) {
			return f
		}
	}
	return 0
}

func main() {
	largestPrimeFactor := LargestPrimeFactor(13195)
	fmt.Println(largestPrimeFactor)
}

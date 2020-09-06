package main

import (
	"fmt"
	"math"
	"time"
)

func sieveOfEratosthenes(n float64) []float64 {
	result := []float64{}
	a := make([]bool, int(n)+1)
	// set a values all to true
	for i, _ := range a {
		a[i] = true
	}

	var i float64
	for i = 2; i <= math.Sqrt(n); i++ {
		x := int(i)
		if a[x] {
			for j := x * x; j <= int(n); j += x {
				a[j] = false
			}
		}
	}

	for i = 2; i <= n; i++ {
		x := int(i)
		if a[x] {
			result = append(result, i)
		}
	}

	return result
}

func sumOfPrimesBelowN(n float64) float64 {
	primes := sieveOfEratosthenes(n)
	var sum float64 = 0

	for _, n := range primes {
		sum += n
	}

	return sum
}

func main() {
	start := time.Now()
	fmt.Println(sumOfPrimesBelowN(2000000))
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"time"
)

func sieveOfEratosthenes(n float64) []float64 {
	var result []float64
	a := make([]bool, int(n)+1)

	// initialize a to all true
	for i, _ := range a {
		a[i] = true
	}

	var i float64
	for i = 2; i <= math.Sqrt(n); i++ {
		x := int(i)
		if a[x] {
			for j := x * x; j < int(n); j += x {
				a[j] = false
			}
		}
	}

	for j := 2; j < int(n); j++ {
		if a[j] {
			result = append(result, float64(j))
		}
	}

	return result
}

func nthPrime(n int) float64 {
	primes := sieveOfEratosthenes(1000000)
	return primes[n-1]
}

func main() {
	start := time.Now()
	fmt.Println(nthPrime(10001))
	fmt.Printf("Solution took %v\n", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"sort"
)

func sieveOfEratosthenes(n int) []int {
	result := []int{}
	a := make([]bool, int(n)+1)

	// initialize a to all true
	for i := range a {
		a[i] = true
	}

	var i int
	limit := int(math.Round(math.Sqrt(float64(n))))
	for i = 2; i <= limit; i++ {
		x := int(i)
		if a[x] {
			for j := x * x; j < int(n); j += x {
				a[j] = false
			}
		}
	}

	for j := 2; j < int(n); j++ {
		if a[j] {
			result = append(result, j)
		}
	}

	return result
}

func Solution() {
	result := 0
	limit := 1000000
	primes := sieveOfEratosthenes(limit)
	primeSum := make([]int, len(primes)+1)

	primeSum[0] = 0
	for i := 0; i < len(primes); i++ {
		primeSum[i+1] = primeSum[i] + primes[i]
	}

	numberOfPrimes := 0
	for i := numberOfPrimes; i < len(primeSum); i++ {
		for j := i - (numberOfPrimes + 1); j >= 0; j-- {
			t := primeSum[i] - primeSum[j]
			if t > limit {
				break
			}

			if s := sort.SearchInts(primes, t); primes[s] == t {
				numberOfPrimes = i - j
				result = t
			}
		}
	}

	fmt.Println(numberOfPrimes, result)
}

func main() {
	Solution()
}

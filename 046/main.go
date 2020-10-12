package main

import (
	"fmt"
	"math"
)

var primes []int

func sieveOfEratosthenes(n int) []int {
	var result []int
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
			result = append(result, int(j))
		}
	}

	return result
}

func Check(n int) bool {
	for _, p := range primes {
		if p > n {
			break
		}

		if p == n {
			return true
		}

		prod := 0
		for j := 1; prod+p < n; j++ {
			prod = 2 * j * j
		}
		if prod+p == n {
			return true
		}
	}
	return false
}

func Solution() {
	for n := 35; n <= 10000; n += 2 {
		c := Check(n)
		if !c {
			fmt.Println(n, c)
		}
	}
}

func main() {
	primes = sieveOfEratosthenes(10000)
	Solution()
}

package main

import (
	"fmt"
	"math"
	"time"
)

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

func GetPrimeFactors(num int) map[int]int {
	factors := make(map[int]int)

	for num%2 == 0 {
		if count, ok := factors[2]; ok {
			factors[2] = count + 1
		} else {
			factors[2] = 1
		}

		num = num / 2
	}

	limit := int(math.Sqrt(float64(num))) + 1
	for i := 3; i <= limit; i += 2 {
		for num%i == 0 {
			if count, ok := factors[i]; ok {
				factors[i] = count + 1
			} else {
				factors[i] = 1
			}
			num = num / i
		}
	}

	if num > 2 {
		factors[num] = 1
	}

	return factors
}

func GetDistinctPrimeFactors(num int) []int {
	f := GetPrimeFactors(num)
	factors := []int{}

	for k := range f {
		factors = append(factors, k)
	}

	return factors
}

func Solution() {
	count := 0
	start := 0
	for i := 100000; i <= 150000; i++ {
		f := GetDistinctPrimeFactors(i)
		if len(f) == 4 {
			start = i
			count++
		} else {
			count = 0
		}

		if count == 4 {
			break
		}
	}

	fmt.Println(start - 3)
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// From problem 7 solution
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

func IsPandigital(num int, l int) bool {
	digits := make([]int, l+1)

	for num > 0 {
		digit := num % 10
		if digit > l || digit == 0 || digits[digit] != 0 {
			return false
		}
		digits[digit] = 1
		num = num / 10
	}

	return true
}

func main() {
	start := time.Now()
	primes := sieveOfEratosthenes(987654321)
	fmt.Printf("Primes took %v", time.Since(start))
	for i := len(primes) - 1; i >= 0; i-- {
		p := primes[i]
		s := strconv.Itoa(p)
		if IsPandigital(p, len(s)) {
			fmt.Println(p)
			break
		}
	}
	fmt.Printf("Solution took %v", time.Since(start))
}

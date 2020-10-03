package main

import (
	"fmt"
	"math"
	"time"
)

var primes []int

// from solution 35
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

// from solution 35
func GetDigits(num int) []int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	return Reverse(digits)
}

func Reverse(digits []int) []int {
	result := make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		result[len(digits)-1-i] = digits[i]
	}
	return result
}

func Pow(a, b int) int {
	total := 1
	for i := 0; i < b; i++ {
		total *= a
	}
	return total
}

func Combine(digits []int) int {
	sum := 0
	l := len(digits)
	for i := 0; i < l; i++ {
		mult := Pow(10, l-i-1)
		sum += digits[i] * mult
	}
	return sum
}

func IsPrime(num int) bool {
	for _, n := range primes {
		if n > num {
			return false
		}
		if n == num {
			return true
		}
	}
	return false
}

func IsTruncatablePrime(digits []int) bool {
	n := 0
	m := Combine(digits)
	for i := len(digits) - 1; i >= 0; i-- {
		n += digits[i] * Pow(10, len(digits)-1-i)
		if !IsPrime(n) {
			return false
		}

		m = m / 10
		if m > 0 && !IsPrime(m) {
			return false
		}
	}

	return true
}

func Solution() int {
	sum := 0
	for _, n := range primes {
		if n > 10 {
			digits := GetDigits(n)
			if IsTruncatablePrime(digits) {
				sum += n
			}
		}
	}
	return sum
}

func main() {
	start := time.Now()
	primes = sieveOfEratosthenes(1000000)
	fmt.Println(Solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

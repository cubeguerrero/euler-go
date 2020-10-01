package main

import (
	"fmt"
	"math"
	"time"
)

func Rotate(digits []int) {
	temp := digits[0]
	for i := 1; i < len(digits); i++ {
		digits[i-1] = digits[i]
	}
	digits[len(digits)-1] = temp
}

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

// From problem 34 solution
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

func IsCircularPrime(num int, primes []int) bool {
	digits := GetDigits(num)

	for i := 1; i < len(digits); i++ {
		Rotate(digits)
		n := Combine(digits)
		if !Contains(n, primes) {
			return false
		}
	}
	return true
}

func Contains(num int, nums []int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
		if n > num {
			return false
		}
	}
	return false
}

func Solution(max int) int {
	nums := []int{}
	primes := sieveOfEratosthenes(max)
	for _, p := range primes {
		if IsCircularPrime(p, primes) {
			nums = append(nums, p)
		}
	}
	return len(nums)
}

func main() {
	start := time.Now()
	fmt.Println(Solution(1000000))
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"time"
)

func numIsPalindrome(n int) bool {
	digits := getDigits(n)
	mid := len(digits) / 2

	for i := 0; i <= mid; i++ {
		a := digits[i]
		b := digits[len(digits)-i-1]
		if a != b {
			return false
		}
	}

	return true
}

func getDigits(n int) []int {
	digits := make([]int, 0)

	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	return digits
}

func largestPalindromeOfNDigitNumbers(n int) int {
	start := int(math.Pow(10, float64(n-1)))
	max := int(math.Pow(10, float64(n))) - 1
	largest := 0

	for i := max; i >= start; i-- {
		for j := max; j >= start; j-- {
			product := i * j
			if largest < product && numIsPalindrome(product) {
				largest = product
			}
		}
	}

	return largest
}

func main() {
	start := time.Now()
	fmt.Println(largestPalindromeOfNDigitNumbers(3))
	fmt.Printf("Solution took %v s\n", time.Since(start))
}

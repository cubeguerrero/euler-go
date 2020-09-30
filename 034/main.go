package main

import (
	"errors"
	"fmt"
	"time"
)

func GetDigits(num int) []int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	return digits
}

func Factorial(num int) int {
	if num < 0 {
		panic(errors.New("Cannot be negative"))
	}

	total := 1
	for i := 1; i <= num; i++ {
		total *= i
	}
	return total
}

func DigitFactorials(num int) int {
	digits := GetDigits(num)
	sum := 0

	for _, n := range digits {
		sum += Factorial(n)
	}

	return sum
}

func Solution() {
	total := 0
	for i := 3; i <= 100000; i++ {
		df := DigitFactorials(i)
		if df == i {
			total += i
		}
	}

	fmt.Println(total)
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

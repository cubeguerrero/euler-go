package main

import (
	"fmt"
	"time"
)

func getDigits(num int) []int {
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}
	return nums
}

func sumOfPower(digits []int, power int) int {
	total := 0
	for _, d := range digits {
		total += pow(d, power)
	}
	return total
}

func pow(a, b int) int {
	total := 1
	for b > 0 {
		total *= a
		b--
	}
	return total
}

func solution() {
	sum := 0
	for i := 2; i <= 354294; i++ {
		digits := getDigits(i)
		if sumOfPower(digits, 5) == i {
			sum += i
		}
	}
	fmt.Println(sum)
}

func main() {
	start := time.Now()
	solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

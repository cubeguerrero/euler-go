package main

import (
	"fmt"
	"time"
)

func Contains(nums []int, n int) bool {
	for _, c := range nums {
		if c == n {
			return true
		}
	}
	return false
}

func getDecimalsOfUnitFraction(n int) ([]int, bool) {
	d := 1
	nums := []int{}
	last := []int{d}
	repeating := false
	for d > 0 {
		d *= 10
		nums = append(nums, d/n)
		d = d % n
		if Contains(last, d) {
			repeating = true
			break
		} else {
			last = append(last, d)
		}
	}

	return nums, repeating
}

func solution(n int) {
	longestRepeatingCount := 0
	longest := 0
	for i := 2; i < n; i++ {
		nums, repeating := getDecimalsOfUnitFraction(i)
		if repeating {
			if len(nums) > longestRepeatingCount {
				longestRepeatingCount = len(nums)
				longest = i
			}
		}
	}
	fmt.Println(longest, longestRepeatingCount)
}

func main() {
	start := time.Now()
	solution(1000)
	fmt.Printf("Solution took %v", time.Since(start))
}

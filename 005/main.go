package main

import (
	"fmt"
	"time"
)

func product(nums []int) int {
	total := 1
	for _, n := range nums {
		total *= n
	}
	return total
}

func leastCommonMultiple(nums []int) int {
	for i := 0; i < len(nums); i++ {
		d := nums[i]
		for j := i + 1; j < len(nums); j++ {
			c := nums[j]
			if c%d == 0 {
				nums[j] = c / d
			}
		}

	}

	return product(nums)
}

func main() {
	start := time.Now()
	fmt.Println(leastCommonMultiple([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))
	fmt.Printf("Solution took: %v\n", time.Since(start))
}

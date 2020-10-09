package main

import (
	"fmt"
	"math"
	"time"
)

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func nextPermutation(nums []int) bool {
	i := len(nums) - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	if i <= 0 {
		return false
	}

	j := len(nums) - 1
	for nums[j] <= nums[i-1] {
		j--
	}

	swap(nums, i-1, j)

	j = len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
	return true
}

func combine(nums []int) int {
	total := 0

	for i := len(nums) - 1; i >= 0; i-- {
		total += nums[i] * int(math.Pow(10, float64(len(nums)-1-i)))
	}

	return total
}

func isSubStringDivisble(nums []int) bool {
	divisors := []int{1, 2, 3, 5, 7, 11, 13, 17}
	for i := 1; i <= len(nums)-3; i++ {
		divisor := divisors[i]
		numerator := combine(nums[i : i+3])
		if numerator%divisor != 0 {
			return false
		}
	}
	return true
}

func solution() int {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := 0

	for {
		ok := nextPermutation(nums)
		if !ok {
			break
		}

		if nums[0] == 0 {
			continue
		}

		if isSubStringDivisble(nums) {
			n := combine(nums)
			total += n
		}
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

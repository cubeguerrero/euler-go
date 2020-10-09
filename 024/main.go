package main

import (
	"errors"
	"fmt"
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

	if i < 0 {
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

func solution(m int) ([]int, error) {
	iterations := 1
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for iterations < m {
		ok := nextPermutation(nums)
		if !ok {
			return nil, errors.New("No permutation for" + string(m))
		}
		iterations++
	}

	return nums, nil
}

func main() {
	start := time.Now()
	val, err := solution(1000000)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	fmt.Printf("Solution took %v", time.Since(start))
}

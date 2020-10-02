package main

import (
	"fmt"
	"strconv"
	"time"
)

func IsPalindrome(input string) bool {
	mid := len(input) / 2
	j := len(input)
	for i := 0; i <= mid; i++ {
		j -= 1
		if input[i] != input[j] {
			return false
		}
	}
	return true
}

func Solution(m int64) int64 {
	var sum int64 = 0
	var i int64
	for i = 1; i < m; i++ {
		base_10 := strconv.FormatInt(i, 10)
		base_2 := strconv.FormatInt(i, 2)

		if IsPalindrome(base_10) && IsPalindrome(base_2) {
			sum += i
		}
	}

	return sum
}

func main() {
	start := time.Now()
	fmt.Println(Solution(1000000))
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

func IsPandigital(num string) bool {
	digits := make(map[rune]int)

	for _, r := range num {
		if r == '0' {
			return false
		}
		if _, ok := digits[r]; ok {
			return false
		}
		digits[r] = 1
	}

	return true
}

func Solution() {
	for i := 9387; i >= 9234; i-- {
		res := strconv.Itoa(i) + strconv.Itoa(2*i)
		if IsPandigital(res) {
			fmt.Println(res)
			return
		}
	}
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

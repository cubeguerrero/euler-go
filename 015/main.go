package main

import (
	"fmt"
	"time"
)

func catalanNumber(n int, memo map[int]int) int {
	if n <= 1 {
		return 1
	}

	res := 0
	for i := 0; i < n; i++ {
		v := 1
		if val, ok := memo[i]; ok {
			v = val
		} else {
			memo[i] = catalanNumber(i, memo)
			v = memo[i]
		}

		if val, ok := memo[n-i-1]; ok {
			v *= val
		} else {
			memo[n-i-1] = catalanNumber(n-i-1, memo)
			v *= memo[n-i-1]
		}
		res += v
	}
	return res
}

func main() {
	start := time.Now()
	memo := make(map[int]int)
	fmt.Println(catalanNumber(20, memo))
	fmt.Printf("Solution took %v", time.Since(start))
}

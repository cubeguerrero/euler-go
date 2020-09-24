package main

import (
	"fmt"
	"time"
)

func solution(n int) int {
	total := 1
	for i := 3; i <= n; i += 2 {
		m := i * i
		j := i - 1
		total += m
		total += m - j
		total += m - (j * 2)
		total += m - (j * 3)
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Println(solution(1001))
	fmt.Printf("Solution took %v", time.Since(start))
}

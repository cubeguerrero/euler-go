package main

import (
	"fmt"
	"time"
)

func GCD(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a > b {
		return GCD(a-b, b)
	}

	return GCD(a, b-a)
}

func solution() int {
	dp := 1
	np := 1

	for c := 1; c <= 9; c++ {
		for d := 1; d < c; d++ {
			for n := 1; n < d; n++ {
				if (n*10+c)*d == (c*10+d)*n {
					np *= n
					dp *= d
				}
			}
		}
	}

	return dp / GCD(np, dp)
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

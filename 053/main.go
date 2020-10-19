package main

import (
	"fmt"
	"time"
)

var memo map[float64]float64

func factorial(n float64) float64 {
	if v, ok := memo[n]; ok {
		return v
	}

	if n <= 0 {
		return 1
	}

	var total float64 = 1
	var i float64
	for i = 2; i <= n; i++ {
		total *= i
	}
	return total
}

func combinations(n, r float64) float64 {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func solution() {
	count := 0
	var n, r float64
	for n = 1; n <= 100; n++ {
		for r = 1; r <= n; r++ {
			c := combinations(n, r)
			if c >= 1000000 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func main() {
	start := time.Now()
	memo = make(map[float64]float64)
	solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

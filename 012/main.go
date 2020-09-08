package main

import (
	"fmt"
	"math"
	"time"
)

func triangularNumber(n int) int {
	return (n * (n + 1)) / 2
}

func getDivisors(n int) []int {
	limit := int(math.Round(math.Sqrt(float64(n)))) + 1
	divisors := []int{}
	for i := 1; i < limit; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

func solution(n int) int {
	i := 1
	for {
		t := triangularNumber(i)
		divisors := getDivisors(t)
		if len(divisors) > n {
			return t
		}
		i++
	}
}

func main() {
	start := time.Now()
	fmt.Println(solution(500))
	fmt.Printf("Solution took %v", time.Since(start))
}

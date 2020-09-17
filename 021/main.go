package main

import (
	"fmt"
	"math"
	"time"
)

func getDivisors(n float64) []float64 {
	divisors := []float64{1}
	var i float64
	for i = 2; i < math.Sqrt(n); i++ {
		if math.Mod(n, i) == 0 {
			divisors = append(divisors, i)
			divisors = append(divisors, n/i)
		}
	}
	return divisors
}

func sum(nums []float64) float64 {
	var t float64 = 0
	for _, n := range nums {
		t += n
	}
	return t
}

func getAmicablePair(a float64) (float64, bool) {
	divisorsA := getDivisors(a)
	b := sum(divisorsA)
	divisorsB := getDivisors(b)
	if sum(divisorsB) == a && b != a {
		return b, true
	} else {
		return -1, false
	}
}

func solution(n float64) float64 {
	memo := make(map[float64]float64)
	var i float64
	for i = 1; i < n; i++ {
		b, found := getAmicablePair(i)
		if found {
			memo[i] = b
			memo[b] = i
		}
	}

	var sum float64 = 0
	for key := range memo {
		sum += key
	}

	return sum
}

func main() {
	start := time.Now()
	fmt.Println(solution(10000))
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"time"
)

type Pair struct {
	A int
	B int
}

func (p Pair) product() int {
	return p.A * p.B
}

func (p Pair) quadraticFormula(n int) int {
	return n*n + (p.A * n) + p.B
}

func isPrime(n int) bool {
	if n < 0 {
		return false
	}

	limit := int(math.Round(math.Pow(float64(n), 0.5)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func solution() int {
	data := make(map[Pair]int)
	longest := 0
	longestPair := Pair{}

	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			p := Pair{A: a, B: b}
			n := 0
			count := 0
			for {
				ans := p.quadraticFormula(n)
				if !isPrime(ans) {
					break
				}
				n++
				count++
			}

			if longest < count {
				longest = count
				longestPair = p
			}

			data[p] = count
		}
	}

	return longestPair.product()
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

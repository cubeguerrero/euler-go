package main

import (
	"fmt"
	"math/big"
	"time"
)

func solution(max int) int {
	set := make(map[string]struct{})

	for a := 2; a <= max; a++ {
		for b := 2; b <= max; b++ {
			c := Pow(a, b)
			if _, ok := set[c]; !ok {
				set[c] = struct{}{}
			}
		}
	}

	return len(set)
}

func Pow(a, b int) string {
	total := big.NewInt(1)
	c := big.NewInt(int64(a))
	for i := 0; i < b; i++ {
		total.Mul(total, c)
	}
	return total.String()
}

func main() {
	start := time.Now()
	fmt.Println(solution(100))
	fmt.Printf("Solution took %v", time.Since(start))
}

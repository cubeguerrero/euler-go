package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func factorial(n int64) *big.Int {
	total := big.NewInt(1)

	for i := n; i >= 1; i-- {
		y := big.NewInt(i)
		total = total.Mul(total, y)
	}

	return total
}

func solution(n int64) int {
	total := 0
	f := factorial(n)
	digits := strings.Split(f.String(), "")
	for _, d := range digits {
		i, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		total += i
	}
	return total
}

func main() {
	start := time.Now()
	fmt.Println(solution(100))
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"math/big"
	"time"
)

func solution() int {
	a := big.NewInt(1)
	b := big.NewInt(1)
	l := len(b.String())
	iteration := 2
	for l < 1000 {
		temp := b
		b = (big.NewInt(0)).Add(a, b)
		a = temp

		l = len(b.String())
		iteration++
	}

	return iteration
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

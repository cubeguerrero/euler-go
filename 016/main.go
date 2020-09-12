package main

import (
	"fmt"
	"math/big"
	"time"
)

func getDigitSum(n *big.Int) *big.Int {
	z := big.NewInt(0)
	ten := big.NewInt(10)
	sum := big.NewInt(0)
	m := big.NewInt(1)
	for n.Cmp(z) != 0 {
		n, m = n.DivMod(n, ten, m)
		sum = sum.Add(sum, m)
	}
	return sum
}

func solution(n int64) int64 {
	v := new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)
	sum := getDigitSum(v)
	return sum.Int64()
}

func main() {
	start := time.Now()
	fmt.Println(solution(1000))
	fmt.Printf("Solution took %v", time.Since(start))
}

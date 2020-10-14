package main

import (
	"fmt"
	"math/big"
	"time"
)

func Solution() {
	total := big.NewInt(0)

	var i int64
	for i = 1; i <= 1000; i++ {
		a := big.NewInt(i)
		a.Exp(a, a, nil)
		total.Add(total, a)
	}

	s := total.String()
	l := len(s)
	fmt.Println(s[l-10 : l])
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

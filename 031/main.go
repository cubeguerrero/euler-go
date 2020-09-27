package main

import (
	"fmt"
	"time"
)

var coins [8]int = [8]int{1, 2, 5, 10, 20, 50, 100, 200}

func coinSum(total int) int {
	ways := make([]int, total+1)
	ways[0] = 1
	for _, coin := range coins {
		for j := coin; j <= total; j++ {
			ways[j] += ways[j-coin]
		}
	}
	return ways[total]
}

func main() {
	start := time.Now()
	fmt.Println(coinSum(200))
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"time"
)

func sumOfMultiplesOf3And5(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	start := time.Now()
	fmt.Println(sumOfMultiplesOf3And5(1000))

	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

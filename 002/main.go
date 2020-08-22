package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	a := 1
	b := 1
	temp := 1
	sum := 0

	for {
		temp = b
		b = a + b
		a = temp

		if b > n {
			break
		}

		if b%2 == 0 {
			sum += b
		}
	}

	return sum
}

func main() {
	start := time.Now()
	answer := fibonacci(4000000)
	fmt.Println("Answer is", answer)

	elasped := time.Since(start)
	fmt.Println("Solution took %s", elasped)
}

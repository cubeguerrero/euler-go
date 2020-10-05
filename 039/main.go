package main

import (
	"fmt"
	"time"
)

func Solution() {
	result := 0
	resultSolutions := 0

	for p := 2; p <= 1000; p += 2 {
		numberOfSolutions := 0
		for a := 2; a < (p / 3); a++ {
			if p*(p-2*a)%(2*(p-a)) == 0 {
				numberOfSolutions++
			}
		}

		if numberOfSolutions > resultSolutions {
			resultSolutions = numberOfSolutions
			result = p
		}
	}

	fmt.Println(result)
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

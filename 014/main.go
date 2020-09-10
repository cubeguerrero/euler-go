package main

import (
	"fmt"
	"time"
)

func collatzSequenceLength(n int) int {
	l := 1
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		l += 1
	}

	return l
}

func solution() {
	currentLongestLength := 0
	currentLongest := 0

	for i := 2; i < 1000000; i++ {
		l := collatzSequenceLength(i)
		if l > currentLongestLength {
			currentLongestLength = l
			currentLongest = i
		}
	}

	fmt.Println(currentLongest, currentLongestLength)
}

func main() {
	start := time.Now()
	solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

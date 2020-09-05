package main

import (
	"fmt"
	"time"
)

func isPythagoreanTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

func specialPythagoreanTriplet() int {
	for a := 1; a <= 500; a++ {
		for b := a + 1; b <= 500; b++ {
			for c := b + 1; c <= 500; c++ {
				sum := a + b + c
				if sum == 1000 && isPythagoreanTriplet(a, b, c) {
					return a * b * c
				}
			}
		}
	}

	return 0
}

func main() {
	start := time.Now()
	fmt.Println(specialPythagoreanTriplet())
	fmt.Printf("Solution took %v", time.Since(start))
}

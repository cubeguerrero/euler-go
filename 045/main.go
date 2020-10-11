package main

import (
	"fmt"
	"math"
	"time"
)

func IsPentagonal(x int) bool {
	test := (1 + math.Sqrt(1+24*float64(x))) / 6
	return float64(int(test)) == test
}

func IsTriangular(x int) bool {
	test := (-1 + math.Sqrt(1+8*float64(x))) / 2
	return float64(int(test)) == test
}

func GetHexagonal(n int) int {
	return n * (2*n - 1)
}

func Solution() int {
	n := 144
	notFound := true
	result := 0

	for notFound {
		h := GetHexagonal(n)
		if IsPentagonal(h) && IsTriangular(h) {
			notFound = false
			result = h
			break
		}
		n++
	}

	return result
}

func main() {
	start := time.Now()
	fmt.Println(Solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

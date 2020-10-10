package main

import (
	"fmt"
	"math"
	"time"
)

func IsPentagonal(x int) bool {
	n := (math.Sqrt(24*float64(x)+1) + 1) / 6
	return float64(int(n)) == n
}

func Pentagonal(n int) int {
	return (n * (3*n - 1)) / 2
}

func Solution() int {
	result := 0
	pentagonalNumbers := []int{1}
	notFound := true
	i := 2

	for notFound {
		nextNum := Pentagonal(i)

		for j := len(pentagonalNumbers) - 1; j >= 0; j-- {
			if IsPentagonal(nextNum-pentagonalNumbers[j]) && IsPentagonal(nextNum+pentagonalNumbers[j]) {
				result = nextNum - pentagonalNumbers[j]
				notFound = false
				break
			}
		}
		pentagonalNumbers = append(pentagonalNumbers, nextNum)
		i++
	}

	return result
}

func main() {
	start := time.Now()
	fmt.Println(Solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

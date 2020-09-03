package main

import (
	"fmt"
	"time"
)

func getNumbers(n int) []int {
	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func sum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}

func _map(numbers []int, f func(int) int) []int {
	result := []int{}
	for _, n := range numbers {
		result = append(result, f(n))
	}
	return result
}

func sumSquareDifference(n int) int {
	numbers := getNumbers(n)
	total := sum(numbers)
	totalSquared := total * total
	squaredTotal := sum(_map(numbers, func(i int) int {
		return i * i
	}))
	return totalSquared - squaredTotal
}

func main() {
	start := time.Now()
	fmt.Println(sumSquareDifference(100))
	fmt.Printf("Solution took - %v\n", time.Since(start))
}

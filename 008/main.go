package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func convertNumbersLine(numbers string) []int {
	var result []int
	nums := strings.Split(numbers, "")
	for _, n := range nums {
		i, err := strconv.Atoi(n)
		must(err)
		result = append(result, i)
	}
	return result
}

func readNumbers() []int {
	var result []int
	fileName := "numbers.txt"
	f, err := os.Open(fileName)
	must(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		nums := convertNumbersLine(line)
		result = append(result, nums...)
	}

	return result
}

func product(numbers []int) int {
	p := 1
	for _, n := range numbers {
		p *= n
	}
	return p
}

func largestProductInASeries(series []int, step int) int {
	largest := 0
	for i := 0; i <= len(series)-step; i++ {
		nums := series[i : i+step]
		p := product(nums)
		if p > largest {
			largest = p
		}
	}
	return largest
}

func main() {
	series := readNumbers()
	start := time.Now()
	fmt.Println(largestProductInASeries(series, 13))
	fmt.Printf("Solution took %v\n", time.Since(start))
}

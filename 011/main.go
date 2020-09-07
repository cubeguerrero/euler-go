package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Came from 080 solution
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Came from 080 solution
func convertNumbersLine(numbers string) []int {
	var result []int
	nums := strings.Split(numbers, " ")
	for _, n := range nums {
		i, err := strconv.Atoi(n)
		must(err)
		result = append(result, i)
	}
	return result
}

// Came from 080 solution with slight difference
func readNumbers() [][]int {
	var result [][]int
	fileName := "numbers.txt"
	f, err := os.Open(fileName)
	must(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		nums := convertNumbersLine(line)
		result = append(result, nums)
	}

	return result
}

func product(a, b, c, d int) int {
	return a * b * c * d
}

func getLargestProduct() int {
	numbers := readNumbers()
	largest := 0
	for col := 0; col < len(numbers); col++ {
		for row := 0; row < len(numbers[col]); row++ {
			// horizontal: row should be 4 less than len(numbers[col])
			if row < len(numbers[col])-4 {
				p := product(numbers[col][row], numbers[col][row+1], numbers[col][row+2], numbers[col][row+3])
				if largest < p {
					largest = p
				}
			}

			// vertical: col should be 4 less than len(numbers)
			if col < len(numbers)-4 {
				p := product(numbers[col][row], numbers[col+1][row], numbers[col+2][row], numbers[col+3][row])
				if largest < p {
					largest = p
				}
			}

			// back diagonal (\): col should be 4 less than len(numbers) and row should be 4 less than len(numbers[col])
			if col < len(numbers)-4 && row < len(numbers[col])-4 {
				p := product(numbers[col][row], numbers[col+1][row+1], numbers[col+2][row+2], numbers[col+3][row+3])
				if largest < p {
					largest = p
				}
			}

			// front diagonal (/): col should be 4 less than len(numbers) and row should be 3 or greater
			if col < len(numbers)-4 && row >= 3 {
				p := product(numbers[col][row], numbers[col+1][row-1], numbers[col+2][row-2], numbers[col+3][row-3])
				if largest < p {
					largest = p
				}
			}
		}
	}

	return largest
}

func main() {
	start := time.Now()
	fmt.Println(getLargestProduct())
	fmt.Printf("Solution took %v\n", time.Since(start))
}

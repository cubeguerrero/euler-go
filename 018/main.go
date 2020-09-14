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

func splitLine(line string) ([]int, error) {
	str := strings.Split(line, " ")
	result := []int{}
	for _, s := range str {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

func readNumbers() [][]int {
	result := [][]int{}
	f, err := os.Open("numbers.txt")
	must(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		nums, err := splitLine(text)
		must(err)
		result = append(result, nums)
	}
	return result
}

func solution(input [][]int) int {
	input_len := len(input)
	for i := input_len - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			a := input[i+1][j]
			b := input[i+1][j+1]
			c := input[i][j]

			if c+a >= c+b {
				input[i][j] = c + a
			} else {
				input[i][j] = c + b
			}
		}
	}
	return input[0][0]
}

func main() {
	input := readNumbers()
	start := time.Now()
	result := solution(input)
	fmt.Println(result)
	fmt.Printf("Solution took %v", time.Since(start))
}

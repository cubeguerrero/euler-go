package main

import (
	"fmt"
	"strconv"
	"time"
)

func IsPandigital(num string) bool {
	digits := make(map[rune]int)

	for _, r := range num {
		if r == '0' {
			return false
		}
		if _, ok := digits[r]; ok {
			return false
		}
		digits[r] = 1
	}

	return true
}

func Solution() int {
	products := make(map[int]int)
	sum := 0

	for i := 0; i < 9; i++ {
		for j := 999; j < 9999; j++ {
			s := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(i*j)
			if len(s) == 9 && IsPandigital(s) {
				p := i * j
				if _, ok := products[p]; !ok {
					products[p] = 1
					sum += p
				}
			}

			if len(s) > 9 {
				break
			}
		}
	}

	for i := 9; i < 99; i++ {
		for j := 99; j < 999; j++ {
			s := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(i*j)
			if len(s) == 9 && IsPandigital(s) {
				p := i * j
				if _, ok := products[p]; !ok {
					products[p] = 1
					sum += p
				}
			}

			if len(s) > 9 {
				break
			}
		}
	}

	return sum
}

func main() {
	start := time.Now()
	fmt.Println(Solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

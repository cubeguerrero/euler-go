package main

import (
	"fmt"
	"time"
)

func solution() int {
	total := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			m := time.Month(month)
			t := time.Date(year, m, 1, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == time.Sunday {
				total += 1
			}
		}
	}
	return total
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

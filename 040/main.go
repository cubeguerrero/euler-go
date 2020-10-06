package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func getConstant(max int) string {
	constant := ""
	i := 1

	for len(constant) < max {
		constant += strconv.Itoa(i)
		i++
	}

	return constant
}

func Solution(max int) {
	c := getConstant(max)
	total := 1
	positions := []int{0, 9, 99, 999, 9999, 99999, 999999}

	for _, p := range positions {
		val, err := strconv.Atoi(string(c[p]))
		if err != nil {
			log.Fatal(err)
		}
		total *= val
	}

	fmt.Println(total)

}

func main() {
	start := time.Now()
	Solution(1000000)
	fmt.Printf("Solution took %v", time.Since(start))
}

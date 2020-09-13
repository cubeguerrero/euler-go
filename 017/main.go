package main

import (
	"fmt"
	"time"
)

var belowTwenty = map[int]string{
	0:  "",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = map[int]string{
	0: "",
	1: "",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func getThousandsLength(n int) int {
	total := 0
	total += len(belowTwenty[n/1000])
	total += len("thousand")
	if i := n % 1000; i > 0 {
		total += getHundredsLength(i)
	}
	return total
}

func getHundredsLength(n int) int {
	total := 0
	total += len(belowTwenty[n/100])
	total += len("hundred")
	if i := n % 100; i > 0 {
		total += len("and")
		total += getTensLength(i)
	}
	return total
}

func getTensLength(n int) int {
	if n < 20 {
		return len(belowTwenty[n])
	} else {
		total := 0
		total += len(tens[n/10])
		total += len(belowTwenty[n%10])
		return total
	}
}

func solution(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		if i >= 1000 {
			total += getThousandsLength(i)
		} else if i >= 100 {
			total += getHundredsLength(i)
		} else {
			total += getTensLength(i)
		}
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Println(solution(1000))
	fmt.Printf("Solution took %v", time.Since(start))
}

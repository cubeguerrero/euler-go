package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func quadraticFormula(a, b, c float64) (float64, float64) {
	x := b * -1
	y := math.Sqrt(math.Pow(b, 2) - (4 * a * c))
	z := 2 * a

	return (x + y) / z, (x - y) / z
}

func readFile() []string {
	fi, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fi)
	scanner.Scan()
	ln := scanner.Text()
	words := strings.Split(ln, ",")

	return words
}

func getWordValue(word string) int {
	total := 0

	for _, r := range word {
		total += int(r) - 64
	}

	return total
}

func isTriangleWord(word string) bool {
	value := getWordValue(word)
	c := value * 2 * -1
	pos, _ := quadraticFormula(1, 1, float64(c))
	return pos == math.Trunc(pos)
}

func Solution() int {
	words := readFile()
	total := 0

	for _, word := range words {
		word = strings.Trim(word, "\"")
		if isTriangleWord(word) {
			total += 1
		}
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Println(Solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var alphabet = map[rune]int{
	'A': 1,
	'B': 2,
	'C': 3,
	'D': 4,
	'E': 5,
	'F': 6,
	'G': 7,
	'H': 8,
	'I': 9,
	'J': 10,
	'K': 11,
	'L': 12,
	'M': 13,
	'N': 14,
	'O': 15,
	'P': 16,
	'Q': 17,
	'R': 18,
	'S': 19,
	'T': 20,
	'U': 21,
	'V': 22,
	'W': 23,
	'X': 24,
	'Y': 25,
	'Z': 26,
}

func getNameScore(name string) int {
	total := 0
	for _, char := range name {
		v := alphabet[char]
		total += v
	}
	return total
}

func readNamesFile() []string {
	names := []string{}
	f, _ := os.Open("names.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		names = strings.Split(scanner.Text(), ",")
	}
	return names
}

func solution(names []string) int {
	sort.Strings(names)

	total := 0
	for i, name := range names {
		total += getNameScore(name) * (i + 1)
	}
	return total
}

func main() {
	names := readNamesFile()
	start := time.Now()
	fmt.Println(solution(names))
	fmt.Printf("Solution took %v", time.Since(start))
}

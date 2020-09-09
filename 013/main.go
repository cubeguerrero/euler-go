package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

func must(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readNumbers() []*big.Int {
	result := []*big.Int{}
	f, err := os.Open("numbers.txt")
	must(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		i, ok := new(big.Int).SetString(text, 10)
		if !ok {
			errorMessage := fmt.Sprintf("Something went wrong when converting %s", text)
			panic(errorMessage)
		}
		result = append(result, i)
	}
	return result
}

func bigIntSum(nums []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, n := range nums {
		sum = sum.Add(sum, n)
	}
	return sum
}

// func getDigits(num *big.Int) []*big.Int {
// 	nums := []*big.Int{}
// 	for num.Cmp(big.NewInt(0)) {

// 	}
// }

func solution() string {
	numbers := readNumbers()
	sum := bigIntSum(numbers)
	digits := strings.Split(sum.String(), "")
	return strings.Join(digits[0:10], "")
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"time"
)

var primes map[int]int

func getFiveDigitPatterns() [][]int {
	retVal := make([][]int, 4)

	retVal[0] = []int{1, 0, 0, 0, 1}
	retVal[1] = []int{0, 1, 0, 0, 1}
	retVal[2] = []int{0, 0, 1, 0, 1}
	retVal[3] = []int{0, 0, 0, 1, 1}

	return retVal
}

func getSixDigitPatterns() [][]int {
	retVal := make([][]int, 10)

	retVal[0] = []int{1, 1, 0, 0, 0, 1}
	retVal[1] = []int{1, 0, 1, 0, 0, 1}
	retVal[2] = []int{1, 0, 0, 1, 0, 1}
	retVal[3] = []int{1, 0, 0, 0, 1, 1}
	retVal[4] = []int{0, 1, 1, 0, 0, 1}
	retVal[5] = []int{0, 1, 0, 1, 0, 1}
	retVal[6] = []int{0, 1, 0, 0, 1, 1}
	retVal[7] = []int{0, 0, 1, 1, 0, 1}
	retVal[8] = []int{0, 0, 1, 0, 1, 1}
	retVal[9] = []int{0, 0, 0, 1, 1, 1}

	return retVal
}

func sieveOfEratosthenes(n int) map[int]int {
	result := make(map[int]int)
	a := make([]bool, int(n)+1)

	// initialize a to all true
	for i := range a {
		a[i] = true
	}

	var i int
	limit := int(math.Round(math.Sqrt(float64(n))))
	for i = 2; i <= limit; i++ {
		x := int(i)
		if a[x] {
			for j := x * x; j < int(n); j += x {
				a[j] = false
			}
		}
	}

	for j := 2; j < int(n); j++ {
		if a[j] {
			result[j] = 1
		}
	}

	return result
}

func isPrime(number int) bool {
	_, ok := primes[number]
	return ok
}

func fillPattern(pattern []int, number int) []int {
	filledPattern := make([]int, len(pattern))
	temp := number

	for i := len(filledPattern) - 1; i >= 0; i-- {
		if pattern[i] == 1 {
			filledPattern[i] = temp % 10
			temp /= 10
		} else {
			filledPattern[i] = -1
		}
	}

	return filledPattern
}

func generateNumber(repNumber int, filledPattern []int) int {
	temp := 0
	for i := 0; i < len(filledPattern); i++ {
		temp = temp * 10
		if filledPattern[i] == -1 {
			temp += repNumber
		} else {
			temp += filledPattern[i]
		}
	}
	return temp
}

func familySize(repNumber int, pattern []int) int {
	familySize := 1
	for i := repNumber + 1; i < 10; i++ {
		if isPrime(generateNumber(i, pattern)) {
			familySize++
		}
	}
	return familySize
}

func Solution() int {
	fiveDigitPatterns := getFiveDigitPatterns()
	sixDigitPatterns := getSixDigitPatterns()

	for i := 11; i < 1000; i += 2 {
		if i%5 == 0 {
			continue
		}

		patterns := [][]int{}
		if i < 100 {
			patterns = fiveDigitPatterns
		} else {
			patterns = sixDigitPatterns
		}

		for j := 0; j < len(patterns); j++ {
			for k := 0; k <= 2; k++ {
				if patterns[j][0] == 0 && k == 0 {
					continue
				}
				pattern := fillPattern(patterns[j], i)
				candidate := generateNumber(k, pattern)
				if candidate < math.MaxInt64 && isPrime(candidate) {
					if familySize(k, pattern) == 8 {
						return candidate
					}
				}
			}
		}
	}
	return 0
}

func main() {
	start := time.Now()
	primes = sieveOfEratosthenes(1000000)
	fmt.Println(Solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

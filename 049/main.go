package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

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

func GetDigits(num int) map[int]int {
	result := make(map[int]int)
	for num > 0 {
		result[num%10] = 1
		num = num / 10
	}

	return result
}

func Solution() {
	primes := sieveOfEratosthenes(10000)

	for a, _ := range primes {
		if a <= 1487 {
			continue
		}
		b := a + 3330
		c := b + 3330

		_, bOk := primes[b]
		_, cOk := primes[c]
		aDigits := GetDigits(a)
		bDigits := GetDigits(b)
		cDigits := GetDigits(c)

		if bOk && cOk && reflect.DeepEqual(aDigits, bDigits) && reflect.DeepEqual(aDigits, cDigits) {
			total := strconv.Itoa(a)
			total += strconv.Itoa(b)
			total += strconv.Itoa(c)
			fmt.Println(total)
			break
		}

	}
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

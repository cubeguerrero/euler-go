package main

import (
	"fmt"
	"time"
)

var limit int = 28123

func PrimeFactors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func Power(p, i int) int {
	result := 1
	for j := 0; j < i; j++ {
		result *= p
	}
	return result
}

func SumOfProperDivisors(n int) int {
	pfs := PrimeFactors(n)

	// key: prime
	// value: prime exponents
	m := make(map[int]int)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	sumOfAllFactors := 1
	for prime, exponents := range m {
		sumOfAllFactors *= (Power(prime, exponents+1) - 1) / (prime - 1)
	}
	return sumOfAllFactors - n
}

func getAbundantNumbers() []int {
	numbers := []int{}
	for i := 1; i <= limit; i++ {
		if SumOfProperDivisors(i) > i {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

func solution() int {
	abundantNums := getAbundantNumbers()

	sumOfAbundantNums := make(map[int]bool)
	for i := 0; i < len(abundantNums); i++ {
		for j := i; j < len(abundantNums); j++ {
			a := abundantNums[i]
			b := abundantNums[j]
			if a+b <= limit {
				sumOfAbundantNums[a+b] = true
			} else {
				break
			}
		}
	}

	sum := 0
	for i := 1; i <= limit; i++ {
		_, ok := sumOfAbundantNums[i]
		if !ok {
			sum += i
		}
	}

	return sum
}

func main() {
	start := time.Now()
	fmt.Println(solution())
	fmt.Printf("Solution took %v", time.Since(start))
}

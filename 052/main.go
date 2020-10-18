package main

import (
	"fmt"
	"reflect"
	"time"
)

func getDigits(num int) map[int]int {
	digits := make(map[int]int)
	for num > 0 {
		d := num % 10
		if _, ok := digits[d]; ok {
			digits[d]++
		} else {
			digits[d] = 1
		}
		num /= 10
	}
	return digits
}

func Solution() {
	for i := 100000; i < 1000000; i++ {
		digits := getDigits(i)
		if reflect.DeepEqual(digits, getDigits(2*i)) && reflect.DeepEqual(digits, getDigits(3*i)) && reflect.DeepEqual(digits, getDigits(4*i)) && reflect.DeepEqual(digits, getDigits(5*i)) && reflect.DeepEqual(digits, getDigits(6*i)) {
			fmt.Println(i)
			break
		}
	}
}

func main() {
	start := time.Now()
	Solution()
	fmt.Printf("Solution took %v", time.Since(start))
}

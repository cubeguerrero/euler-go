package main

import "testing"

func TestSieveOfEratosthenes(t *testing.T) {
	want := []float64{2, 3, 5, 7}
	got := sieveOfEratosthenes(10)
	for _, n := range got {
		if !inSlice(want, n) {
			t.Errorf("sieveOfEratosthenes(%#v) should not have %f", 10, n)
		}
	}
}

func TestSumOfPrimesBelowN(t *testing.T) {
	got := sumOfPrimesBelowN(10)
	var want float64 = 17
	if got != want {
		t.Errorf("sumOfPrimesBelowN(%#v) - got %f - want %f", 10, got, want)
	}
}

func inSlice(nums []float64, n float64) bool {
	for _, x := range nums {
		if x == n {
			return true
		}
	}
	return false
}

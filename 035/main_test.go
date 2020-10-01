package main

import "testing"

func TestIsCircularPrime(t *testing.T) {
	var input int = 197
	primes := sieveOfEratosthenes(1000)
	if !IsCircularPrime(input, primes) {
		t.Errorf("%#v should be a circular prime.", input)
	}
}

func TestRotate(t *testing.T) {
	input := []int{1, 9, 7}
	Rotate(input)
	want := []int{9, 7, 1}

	for i, n := range want {
		if input[i] != n {
			t.Errorf("Position %d in %#v should be %d", i, input, n)
		}
	}
}

func TestCombine(t *testing.T) {
	input := []int{9, 7, 1}
	got := Combine(input)
	want := 971
	if got != want {
		t.Errorf("Combine(%#v) got: %d want: %d", input, got, want)
	}
}

func TestSolution(t *testing.T) {
	input := 100
	want := 13
	got := Solution(input)

	if got != want {
		t.Errorf("Solution(%#v) got: %d want: %d", input, got, want)
	}
}

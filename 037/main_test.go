package main

import "testing"

func TestIsTruncatablePrime(t *testing.T) {
	primes = sieveOfEratosthenes(5000)
	input := []int{3, 7, 9, 7}
	if !IsTruncatablePrime(input) {
		t.Errorf("IsTruncatablePrime(%#v) should be true", input)
	}
}

func TestIsPrime(t *testing.T) {
	primes = sieveOfEratosthenes(5000)
	input := 3797
	if !IsPrime(input) {
		t.Errorf("%#v is a prime", input)
	}

	input = 3798
	if IsPrime(input) {
		t.Errorf("%#v is not a prime", input)
	}
}

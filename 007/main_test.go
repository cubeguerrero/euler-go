package main

import "testing"

func TestSieveOfEratosthenes(t *testing.T) {
	got := sieveOfEratosthenes(15)
	want := []float64{2, 3, 5, 7, 11, 13}
	for i, n := range want {
		if got[i] != n {
			t.Errorf("The %d element of sieveOfEratosthenes(%#v) = %f, want: %f", i, 15, got[i], n)
		}
	}
}

func TestNthPrime(t *testing.T) {
	n := 6
	var want float64 = 13
	got := nthPrime(n)

	if got != want {
		t.Errorf("The %dth prime - got: %f, want: %f", n, got, want)
	}
}

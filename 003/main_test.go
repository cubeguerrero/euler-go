package main

import "testing"

func test_largest_prime_factor_13195(t *testing.T) {
	expected := 29.0
	got := LargestPrimeFactor(13195)
	if got != expected {
		t.Errorf("LargestPrimeFactor(%#v) = %f, expected %f", 13195, got, expected)
	}
}

func test_largest_prime_factor_600851475143(t *testing.T) {
	expected := 6857.0
	got := LargestPrimeFactor(600851475143)
	if got != expected {
		t.Errorf("LargestPrimeFactor(%#v) = %f, expected %f", 600851475143, got, expected)
	}
}

package main

import "testing"

func TestProduct(t *testing.T) {
	input := []int{1, 2, 3, 2, 5, 1, 7, 2, 3, 1}
	want := 2520
	got := product(input)
	if got != want {
		t.Errorf("product(%#v) = %d, want: %d", input, got, want)
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := 2520
	got := leastCommonMultiple(input)
	if got != want {
		t.Errorf("leastCommonMultiple(%#v) = %d, want: %d", input, got, want)
	}
}

package main

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	want := big.NewInt(3628800)
	got := factorial(10)
	if got.Cmp(want) != 0 {
		t.Errorf("factorial(%#v) got: %v, want: %v", 10, got, want)
	}
}

func TestSolution(t *testing.T) {
	want := 27
	got := solution(10)
	if got != want {
		t.Errorf("solution(%#v) got: %d, want: %d", 10, got, want)
	}
}

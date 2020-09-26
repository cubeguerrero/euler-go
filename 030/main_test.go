package main

import "testing"

func TestGetDigits(t *testing.T) {
	input := 123
	got := getDigits(input)
	want := []int{3, 2, 1}

	if !Equal(got, want) {
		t.Errorf("getDigits(%#v) got: %v want: %v", input, got, want)
	}
}

func TestSumOfPower(t *testing.T) {
	digits := []int{1, 6, 3, 4}
	power := 4
	got := sumOfPower(digits, power)
	want := 1634

	if got != want {
		t.Errorf("sumOfPower(%#v, %#v) got: %d want: %d", digits, power, got, want)
	}
}

func TestPow(t *testing.T) {
	num := 2
	power := 4
	got := pow(2, 4)
	want := 16
	if got != want {
		t.Errorf("pow(%#v, %#v) got: %d want: %d", num, power, got, want)
	}
}

func BenchmarkSolution(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution()
	}
}

func BenchmarkSolutionNaiveConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutionNaiveConcurrent()
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := range a {
		if j != b[i] {
			return false
		}
	}

	return true
}

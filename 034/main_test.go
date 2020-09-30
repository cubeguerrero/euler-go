package main

import "testing"

func TestGetDigits(t *testing.T) {
	input := 145
	got := GetDigits(input)
	want := []int{5, 4, 1}

	if !equal(got, want) {
		t.Errorf("GetDigits(%#v) got: %v want: %v", input, got, want)
	}
}

func TestFactorial(t *testing.T) {
	input := 5
	got := Factorial(input)
	want := 120

	if got != want {
		t.Errorf("Factorial(%#v) got: %d want: %d", input, got, want)
	}

	input = 0
	got = Factorial(input)
	want = 1
	if got != want {
		t.Errorf("Factorial(%#v) got: %d want: %d", input, got, want)
	}

	input = -1
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Factorial(%#v) should should panic", input)
		}
	}()
	Factorial(input)
}

func TestDigitFactorials(t *testing.T) {
	input := 145
	got := DigitFactorials(input)
	want := 145
	if got != want {
		t.Errorf("DigitFactorials(%#v) got: %d want: %d", input, got, want)
	}
}

func equal(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i, n := range got {
		if n != want[i] {
			return false
		}
	}

	return true
}

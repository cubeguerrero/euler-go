package main

import "testing"

func TestCombine(t *testing.T) {
	input := []int{1, 2, 3}
	got := combine(input)
	want := 123

	if got != want {
		t.Errorf("combine(%#v) got: %d want: %d", input, got, want)
	}

	input = []int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}
	got = combine(input)
	want = 1406357289

	if got != want {
		t.Errorf("combine(%#v) got: %d want: %d", input, got, want)
	}
}

func TestIsSubStringDivisble(t *testing.T) {
	input := []int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}

	if !isSubStringDivisble(input) {
		t.Errorf("%#v should be isSubStringDivisble", input)
	}
}

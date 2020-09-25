package main

import "testing"

func TestSolution(t *testing.T) {
	input := 5
	want := 15
	got := solution(input)
	if got != want {
		t.Errorf("solution(%#v) got: %d want %d", input, got, want)
	}
}

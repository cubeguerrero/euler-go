package main

import "testing"

func TestIsPentagonal(t *testing.T) {
	input := 40755
	if !IsPentagonal(input) {
		t.Errorf("%#v should be pentagonal", input)
	}
}

func TestIsTriangular(t *testing.T) {
	input := 40755
	if !IsTriangular(input) {
		t.Errorf("%#v should be triangular", input)
	}
}

func TestGetHexagonal(t *testing.T) {
	input := 2
	got := GetHexagonal(input)
	want := 6

	if got != want {
		t.Errorf("GetHexagonal(%#v) got: %d want: %d", input, got, want)
	}

	input = 143
	got = GetHexagonal(input)
	want = 40755

	if got != want {
		t.Errorf("GetHexagonal(%#v) got: %d want: %d", input, got, want)
	}
}

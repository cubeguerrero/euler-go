package main

import "testing"

func TestQuadraticFormula(t *testing.T) {
	var a float64 = 1
	var b float64 = 1
	var c float64 = -110
	got1, got2 := quadraticFormula(a, b, c)

	if got1 != 10 || got2 != -11 {
		t.Errorf("quadraticFormula(%v, %v, %v) got: %f %f, want: %f %f", a, b, c, got1, got2, 10.0, -11.0)
	}
}

func TestGetWordValue(t *testing.T) {
	input := "SKY"
	got := getWordValue(input)
	want := 55

	if got != want {
		t.Errorf("getWordValue(%#v) got: %d want: %d", input, got, want)
	}
}

func TestIsTriangleWord(t *testing.T) {
	input := "SKY"
	ok := isTriangleWord(input)

	if !ok {
		t.Errorf("isTriangleWord(%#v) should be true", input)
	}
}

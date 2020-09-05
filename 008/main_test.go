package main

import "testing"

func TestConvertNumbersLine(t *testing.T) {
	want := []int{1, 2, 3, 4}
	got := convertNumbersLine("1234")
	for i, g := range got {
		if g != want[i] {
			t.Errorf("The %dth number of convertNumbersLine(%#v), want: %d, got: %d", i, "1234", g, want[i])
		}
	}
}

func TestReadNumbers(t *testing.T) {
	got := readNumbers()
	if len(got) != 1000 {
		t.Errorf("Wrong number of items. Want: 1000, Got: %d", len(got))
	}
}

func TestProduct(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := 24
	got := product([]int{1, 2, 3, 4})
	if got != want {
		t.Errorf("product(%#v) - got: %d, want: %d", input, got, want)
	}
}

func TestLargestProductInASeries(t *testing.T) {
	input := []int{1, 4, 3, 2}
	step := 2
	want := 12
	got := largestProductInASeries(input, step)
	if got != want {
		t.Errorf("largestProductInASeries(%#v, %#v) - got: %d, want: %d", input, step, got, want)
	}
}

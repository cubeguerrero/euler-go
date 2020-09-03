package main

import "testing"

func TestGetNumbers(t *testing.T) {
	result := getNumbers(10)
	if len(result) != 10 {
		t.Errorf("getNumbers(%#v) length is %d, want: %d", 10, 10, 10)
	}
}

func TestSum(t *testing.T) {
	input := []int{1, 2, 3}
	want := 6
	got := sum(input)
	if got != want {
		t.Errorf("sum(%#v) = %d, want: %d", input, got, want)
	}
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	result := _map(input, func(i int) int {
		return i * 2
	})

	for i, n := range input {
		if n*2 != result[i] {
			t.Errorf("func was not applied to %d element of %#v, got %d, want %d", i, input, result[i], n*2)
		}
	}
}

func TestSumSquareDifference(t *testing.T) {
	n := 10
	want := 2640
	got := sumSquareDifference(n)
	if got != want {
		t.Errorf("sumSquareDifference(%#v) = %d, want: %d", n, got, want)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestGetPrimeFactors(t *testing.T) {
	input := 644
	want := map[int]int{
		2:  2,
		7:  1,
		23: 1,
	}
	got := GetPrimeFactors(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPrimeFactors(%#v) got: %#v want: %#v", input, got, want)
	}

	input = 645
	want = map[int]int{
		3:  1,
		5:  1,
		43: 1,
	}
	got = GetPrimeFactors(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPrimeFactors(%#v) got: %#v want: %#v", input, got, want)
	}

	input = 646
	want = map[int]int{
		2:  1,
		17: 1,
		19: 1,
	}
	got = GetPrimeFactors(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetPrimeFactors(%#v) got: %#v want: %#v", input, got, want)
	}
}

func TestGetDistinctPrimeFactors(t *testing.T) {
	input := 644
	checker := map[int]int{
		2:  1,
		7:  1,
		23: 1,
	}
	got := GetDistinctPrimeFactors(input)

	for _, n := range got {
		if _, ok := checker[n]; !ok {
			t.Errorf("GetDistinctPrimeFactors(%#v) should include %d", input, n)
		}
	}
}

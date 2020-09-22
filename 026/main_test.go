package main

import "testing"

func TestGetDecimals(t *testing.T) {
	want := []int{5}
	got, repeating := getDecimalsOfUnitFraction(2)

	if repeating {
		t.Errorf("getDecimalsOfUnitFraction(%#v) repeating should not be repeating", 2)
	}

	if !Equal(got, want) {
		t.Errorf("getDecimalsOfUnitFraction(%#v) got: %v want: %v", 2, got, want)
	}

	want = []int{2, 5}
	got, repeating = getDecimalsOfUnitFraction(4)

	if repeating {
		t.Errorf("getDecimalsOfUnitFraction(%#v) repeating should not be repeating", 4)
	}

	if !Equal(got, want) {
		t.Errorf("getDecimalsOfUnitFraction(%#v) got: %v want: %v", 4, got, want)
	}

	want = []int{3}
	got, repeating = getDecimalsOfUnitFraction(3)

	if !repeating {
		t.Errorf("getDecimalsOfUnitFraction(%#v) repeating should be repeating", 3)
	}

	if !Equal(got, want) {
		t.Errorf("getDecimalsOfUnitFraction(%#v) got: %v want: %v", 3, got, want)
	}

	want = []int{1, 4, 2, 8, 5, 7}
	got, repeating = getDecimalsOfUnitFraction(7)

	if !repeating {
		t.Errorf("getDecimalsOfUnitFraction(%#v) repeating should be repeating", 7)
	}

	if !Equal(got, want) {
		t.Errorf("getDecimalsOfUnitFraction(%#v) got: %v want: %v", 7, got, want)
	}
}

func Equal(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i, n := range got {
		if want[i] != n {
			return false
		}
	}

	return true
}

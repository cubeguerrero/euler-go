package main

import "testing"

func TestIsPythagoreanTriplet(t *testing.T) {
	got := isPythagoreanTriplet(3, 4, 5)
	if !got {
		t.Errorf("isPythagoreanTriplet(%v, %v, %v) should be true.\n", 3, 4, 5)
	}
}

func TestSpecialPythagoreanTriplet(t *testing.T) {
	want := 31875000
	got := specialPythagoreanTriplet()
	if got != want {
		t.Errorf("specialPythagoreanTriplet() should be %d, got %d", want, got)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestGetDivisors(t *testing.T) {
	want := []float64{1, 2, 142, 4, 71}
	got := getDivisors(284)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("getDivisors(%#v) got: %v, want: %v", 280, got, want)
	}
}

func TestGetAmicablePairs(t *testing.T) {
	g, b := getAmicablePair(220)
	var w float64 = 284
	if !b {
		t.Errorf("getAmicablePair(%#v) got: %f, want: %f", 220, g, w)
	}
}

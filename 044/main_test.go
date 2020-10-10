package main

import "testing"

func TestIsPentagonal(t *testing.T) {
	input := 35
	if !IsPentagonal(input) {
		t.Errorf("IsPentagonal(%#v) should be pentagonal", input)
	}
}

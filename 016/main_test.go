package main

import "testing"

func TestSolution(t *testing.T) {
	if solution(0) != 1 {
		t.Errorf("solution(%#v) should be 1", 0)
	}

	if solution(15) != 26 {
		t.Errorf("solution(%#v) should be 26", 15)
	}
}

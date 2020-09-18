package main

import "testing"

func TestGetNameScore(t *testing.T) {
	input := "COLIN"
	got := getNameScore(input)
	want := 53

	if got != want {
		t.Errorf("getNameScore(%#v) got: %d want: %d", input, got, want)
	}
}

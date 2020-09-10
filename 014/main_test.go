package main

import "testing"

func TestCollatzSequenceLength(t *testing.T) {
	want := 10
	got := collatzSequenceLength(13)
	if got != want {
		t.Errorf("collatzSequenceLength(%#v) got: %d want: %d", 13, got, want)
	}
}

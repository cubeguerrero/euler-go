package main

import "testing"

func TestNumIsPlaindrome(t *testing.T) {
	if numIsPalindrome(100) {
		t.Errorf("numIsPalindrome(%#v) should be false", 100)
	}

	if !numIsPalindrome(101) {
		t.Errorf("numIsPalindrome(%#v) should be true", 101)
	}
}

func TestGetDigits(t *testing.T) {
	want := []int{0, 0, 1}
	got := getDigits(100)
	if !compareSlice(got, want) {
		t.Errorf("getDigits(%#v) = %v, want: %v", 100, got, want)
	}
}

func TestLargestPalindromeOfNDigitNumbers(t *testing.T) {
	want := 906609
	got := largestPalindromeOfNDigitNumbers(3)
	if got != want {
		t.Errorf("largestPalindromeOfNDigitNumbers(%#v) = %v, want: %v", 3, got, want)
	}
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	inputs := map[string]bool{
		"1001":  true,
		"585":   true,
		"400":   false,
		"100":   false,
		"10101": true,
	}

	for input, result := range inputs {
		if IsPalindrome(input) != result {
			t.Errorf("%v should be a palindrome.\n", input)
		}
	}
}

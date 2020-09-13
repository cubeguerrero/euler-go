package main

import "testing"

func TestSolution(t *testing.T) {
	want := 19
	got := solution(5)

	if got != want {
		t.Errorf("solution(%#v) got: %d, want: %d", 5, got, want)
	}

	want = 21124
	got = solution(1000)

	if got != want {
		t.Errorf("solution(%#v) got: %d, want: %d", 1000, got, want)
	}
}

func TestGetTensLength(t *testing.T) {
	want := 9
	got := getTensLength(21)
	if got != want {
		t.Errorf("getTensLength(%#v) got: %d want: %d", 21, got, want)
	}

	want = 6
	got = getTensLength(20)
	if got != want {
		t.Errorf("getTensLength(%#v) got: %d want: %d", 20, got, want)
	}
}

func TestGetHundredsLength(t *testing.T) {
	want := 23
	got := getHundredsLength(342)
	if got != want {
		t.Errorf("getHundredsLength(%#v) got: %d want: %d", 342, got, want)
	}

	want = 20
	got = getHundredsLength(115)
	if got != want {
		t.Errorf("getHundredsLength(%#v) got: %d want: %d", 115, got, want)
	}

	want = 19
	got = getHundredsLength(111)
	if got != want {
		t.Errorf("getHundredsLength(%#v) got: %d want: %d", 111, got, want)
	}
}

func TestGetThousandsLength(t *testing.T) {
	want := 11
	got := getThousandsLength(1000)
	if got != want {
		t.Errorf("getThousandsLength(%#v) got: %d want: %d", 1000, got, want)
	}
}

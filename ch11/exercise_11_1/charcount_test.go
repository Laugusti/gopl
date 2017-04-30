package charcount

import (
	"testing"
	"unicode/utf8"
)

func TestRuneCount(t *testing.T) {
	var tests = []struct {
		input     string
		runeCount map[rune]int
		utflen    [utf8.UTFMax + 1]int
		invalid   int
	}{
		{"abc", map[rune]int{'a': 1, 'b': 1, 'c': 1}, [utf8.UTFMax + 1]int{1: 3}, 0},
		{"世界", map[rune]int{'世': 1, '界': 1}, [utf8.UTFMax + 1]int{3: 2}, 0},
		{"\xe2\xed\xa1", map[rune]int{}, [utf8.UTFMax + 1]int{}, 3},
		{"\xe2\xed\xa1abc世界", map[rune]int{'a': 1, 'b': 1, 'c': 1, '世': 1, '界': 1}, [utf8.UTFMax + 1]int{1: 3, 3: 2}, 3},
	}
	for _, test := range tests {
		counts, utflen, invalid, err := RuneCount(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if !mapEqual(counts, test.runeCount) {
			t.Errorf("RuneCount (counts) => input: %q, expected: %v, got: %v",
				test.input, test.runeCount, counts)
		}
		if utflen != test.utflen {
			t.Errorf("RuneCount (utflen) => input: %q, expected: %v, got: %v",
				test.input, test.utflen, utflen)
		}
		if invalid != test.invalid {
			t.Errorf("RuneCount (invalid) => input: %q, expected: %v, got: %v",
				test.input, test.invalid, invalid)
		}
	}
}

func mapEqual(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

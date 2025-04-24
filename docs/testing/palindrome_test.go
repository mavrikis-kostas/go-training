package stringutils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"single character", "x", "x"},
        {"palindrome", "aba", "aba"},
        {"simple string", "qwe", "ewq"},
        {"with spaces", "qwe rt", "tr ewq"},
		{"with numbers", "12345", "54321"},
		{"with special characters", "!@#$%^&*()", ")(*&^%$#@!"},
		{"with mixed case", "Hello World", "dlroW olleH"},
		{"with unicode", "こんにちは", "はちにんこ"},
		{"with emojis", "😀😃😄", "😄😃😀"},
    }

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Reverse(tc.input)
			if actual != tc.expected {
				t.Errorf("Reverse(%q) = %q; expected %q", tc.input, actual, tc.expected)
			}
		})
	}
}
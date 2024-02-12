package main

import (
	"testing"
)

func TestBuildString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", `qwe45`},
		{"qwe\\45", `qwe44444`},
		{"qwe\\\\5", `qwe\\\\\`},
	}

	for _, test := range tests {
		result := UnpackingString(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}

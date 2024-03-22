package parser

import (
	"testing"
)

func TestIntToHex(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{16, "0x10"},
		{32, "0x20"},
		{256, "0x100"},
	}

	for _, tc := range tests {
		result := intToHex(tc.input)
		if result != tc.expected {
			t.Errorf("intToHex(%d) - Expected: %s, Got: %s", tc.input, tc.expected, result)
		}
	}
}

func TestHexToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0x10", 16},
		{"0x20", 32},
		{"0x100", 256},
	}

	for _, tc := range tests {
		result := hexToInt(tc.input)
		if result != tc.expected {
			t.Errorf("hexToInt(%s) - Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}

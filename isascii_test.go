package strutil

import (
	"testing"
)

func TestIsASCII(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcdef", true},
		{"    ", true},
		{" '.'._  --  ", true},
		{"", true},
		{"Υπάρχουν", false},
	}

	for i, test := range tests {
		output := IsASCII(test.input)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

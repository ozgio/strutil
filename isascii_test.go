package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

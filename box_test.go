package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBox(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		align    string
		expected string
	}{
		{"Hello World", 20, AlignTypeCenter, "┌──────────────────┐\n│   Hello World    │\n└──────────────────┘"},
		{"\nHello World\n", 20, AlignTypeCenter, "┌──────────────────┐\n│                  │\n│   Hello World    │\n│                  │\n└──────────────────┘"},
		{"résumé", 10, AlignTypeLeft, "┌────────┐\n│résumé  │\n└────────┘"},
	}

	for i, test := range tests {
		output, _ := Box(test.input, test.width, test.align)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

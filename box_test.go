package strutil

import (
	"fmt"
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

func ExampleBox() {
	output, _ := Box("Hello World", 20, AlignTypeCenter)
	fmt.Println(output)
	// Outputs:
	//┌──────────────────┐
	//│   Hello World    │
	//└──────────────────┘
}

func ExampleBox_Long() {
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`
	output, _ := Box(text, 30, AlignTypeLeft)
	fmt.Println(output)
	// Outputs:
	// ┌────────────────────────────┐
	// │Lorem ipsum dolor sit amet, │
	// │consectetur adipiscing elit,│
	// │sed do eiusmod tempor       │
	// │incididunt ut labore et     │
	// │dolore magna aliqua. Ut enim│
	// │ad minim veniam, quis       │
	// │nostrud exercitation ullamco│
	// │laboris nisi ut aliquip ex  │
	// │ea commodo consequat.       │
	// └────────────────────────────┘
}

func ExampleCustomBox() {
	output, _ := CustomBox("Hello World", 20, AlignTypeCenter, DefaultBox9Slice)
	fmt.Println(output)
	// Outputs:
	//┌──────────────────┐
	//│   Hello World    │
	//└──────────────────┘
}

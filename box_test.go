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
		err      bool
	}{
		{"Hello World", 20, AlignTypeCenter, "┌──────────────────┐\n│   Hello World    │\n└──────────────────┘", false},
		{"\nHello World\n", 20, AlignTypeCenter, "┌──────────────────┐\n│                  │\n│   Hello World    │\n│                  │\n└──────────────────┘", false},
		{"résumé", 10, AlignTypeLeft, "┌────────┐\n│résumé  │\n└────────┘", false},
		{"Hello World", 2, AlignTypeLeft, "", true},
		{"Hello\n\n\nWorld", 10, AlignTypeLeft, "┌────────┐\n│Hello   │\n│        │\n│        │\n│World   │\n└────────┘", false},
	}

	for i, test := range tests {
		output, err := Box(test.input, test.width, test.align)
		if test.err {
			assert.Errorf(t, err, "Test case %d is not successful\n", i)
		} else {
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}
}

func ExampleBox() {
	output, _ := Box("Hello World", 20, AlignTypeCenter)
	fmt.Println(output)
	// Output:
	//┌──────────────────┐
	//│   Hello World    │
	//└──────────────────┘
}

func ExampleBox_long() {
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`
	output, _ := Box(text, 30, AlignTypeLeft)
	fmt.Println(output)
	// Output:
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
	// Output:
	//┌──────────────────┐
	//│   Hello World    │
	//└──────────────────┘
}

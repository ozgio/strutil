package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawBox(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		align    string
		expected string
		err      bool
	}{
		{"", 20, AlignTypeCenter, "┌──────────────────┐\n│                  │\n└──────────────────┘", false},
		{"Hello World", 20, AlignTypeCenter, "┌──────────────────┐\n│   Hello World    │\n└──────────────────┘", false},
		{"\nHello World\n", 20, AlignTypeCenter, "┌──────────────────┐\n│                  │\n│   Hello World    │\n│                  │\n└──────────────────┘", false},
		{"résumé", 10, AlignTypeLeft, "┌────────┐\n│résumé  │\n└────────┘", false},
		{"Hello World", 2, AlignTypeLeft, "", true},
		{"Hello\n\n\nWorld", 10, AlignTypeLeft, "┌────────┐\n│Hello   │\n│        │\n│        │\n│World   │\n└────────┘", false},
	}

	for i, test := range tests {
		output, err := DrawBox(test.input, test.width, test.align)
		if test.err {
			assert.Errorf(t, err, "Test case %d is not successful\n", i)
		} else {
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}
}

func TestDrawCustomBox(t *testing.T) {
	test9Slice := SimpleBox9Slice()
	test9Slice.Top = "-৹_৹"
	test9Slice.Left = "Ͼ|"
	test9Slice.Right = "|Ͽ"
	test9Slice.Bottom = "৹-৹"

	tests := []struct {
		input    string
		width    int
		align    string
		expected string
		err      bool
	}{
		{"", 20, AlignTypeCenter, "+-৹_৹-৹_৹-৹_৹-৹_৹-৹+\nϾ|                |Ͽ\n+৹-৹৹-৹৹-৹৹-৹৹-৹৹-৹+", false},
		{"Hello World", 20, AlignTypeCenter, "+-৹_৹-৹_৹-৹_৹-৹_৹-৹+\nϾ|  Hello World   |Ͽ\n+৹-৹৹-৹৹-৹৹-৹৹-৹৹-৹+", false},
		{"\nHello World\n", 20, AlignTypeCenter, "+-৹_৹-৹_৹-৹_৹-৹_৹-৹+\nϾ|                |Ͽ\nϾ|  Hello World   |Ͽ\nϾ|                |Ͽ\n+৹-৹৹-৹৹-৹৹-৹৹-৹৹-৹+", false},
		{"résumé", 10, AlignTypeLeft, "+-৹_৹-৹_৹+\nϾ|résumé|Ͽ\n+৹-৹৹-৹৹-+", false},
		{"résumé", 9, AlignTypeLeft, "+-৹_৹-৹_+\nϾ|résum|Ͽ\nϾ|é    |Ͽ\n+৹-৹৹-৹৹+", false},
		{"Hello World", 2, AlignTypeLeft, "", true},
		{"Hello\n\n\nWorld", 10, AlignTypeRight, "+-৹_৹-৹_৹+\nϾ| Hello|Ͽ\nϾ|      |Ͽ\nϾ|      |Ͽ\nϾ| World|Ͽ\n+৹-৹৹-৹৹-+", false},
	}

	for i, test := range tests {
		output, err := DrawCustomBox(test.input, test.width, test.align, test9Slice, "\n")
		if test.err {
			assert.Errorf(t, err, "Test case %d is not successful, expecting error\n", i)
		} else {
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}
}

func ExampleDrawBox() {
	output, _ := DrawBox("Hello World", 20, AlignTypeCenter)
	fmt.Println(output)
	// Output:
	//┌──────────────────┐
	//│   Hello World    │
	//└──────────────────┘
}

func ExampleDrawBox_long() {
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`
	output, _ := DrawBox(text, 30, AlignTypeLeft)
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

func ExampleDrawCustomBox() {
	output, _ := DrawCustomBox("Hello World", 20, AlignTypeCenter, DefaultBox9Slice(), "\n")
	fmt.Println(output)
	// Output:
	//┌──────────────────┐
	//│   Hello World    │
	//└──────────────────┘
}

func TestTile(t *testing.T) {
	tests := []struct {
		pattern  string
		length   int
		expected string
	}{
		{"", 10, ""},
		{"-", 10, "----------"},
		{"-.", -1, ""},
		{"-.", 0, ""},
		{"-.", 2, "-."},
		{"-.", 1, "-"},
		{"-.", 3, "-.-"},
		{"-.", 4, "-.-."},
		{"-৹", 0, ""},
		{"-৹", 2, "-৹"},
		{"-৹", 1, "-"},
		{"-৹", 3, "-৹-"},
		{"-৹", 4, "-৹-৹"},
	}

	for i, test := range tests {
		output := Tile(test.pattern, test.length)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

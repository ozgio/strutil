package strutil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlign(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		typ      AlignType
		expected string
	}{
		{"  lorem  ", 10, Left, "lorem  "},
		{"  lorem  ", 10, Right, "     lorem"},
		{"  lorem  ", 10, Center, "  lorem   "},
		{"  lorem  ", 10, "", "  lorem  "},
	}

	for i, test := range tests {
		output := Align(test.input, test.typ, test.width)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleAlign() {
	fmt.Println(Align("  lorem  \n  ipsum  ", Right, 10))
	// Output:
	//      lorem
	//      ipsum
}

func TestAlignLeft(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"    lorem", "lorem"},
		{"   lorem\n    ipsum", "lorem\nipsum"},
		{"  lorem  \n  ipsum  \n", "lorem  \nipsum  \n"},
	}

	for i, test := range tests {
		output := AlignLeft(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleAlignLeft() {
	fmt.Println(AlignLeft("   lorem\n    ipsum"))
	// Output:
	// lorem
	// ipsum
}

func TestAlignRight(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		expected string
	}{
		{"    lorem", 10, "     lorem"},
		{"   lorem\n    ipsum", 10, "     lorem\n     ipsum"},
		{"  lorem  \n  ipsum  \n", 10, "     lorem\n     ipsum\n          "},
		{"  lorem  \n  ipsum  \n", 1, "lorem\nipsum\n "},
	}

	for i, test := range tests {
		output := AlignRight(test.input, test.width)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleAlignRight() {
	fmt.Println(AlignRight("  lorem  \n  ipsum  ", 10))
	// Output:
	//      lorem
	//      ipsum
}

func TestAlignCenter(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		expected string
	}{
		{"", 10, "          "},
		{"lorem", 10, "  lorem   "},
		{"lorem\nipsum", 10, "  lorem   \n  ipsum   "},
		{"    lorem", 10, "  lorem   "},
		{"   lorem\n    ipsum", 10, "  lorem   \n  ipsum   "},
		{"  lorem  \n  ipsum  \n", 10, "  lorem   \n  ipsum   \n          "},
		{"  lorem  \n  ipsum  \n", 1, "lorem\nipsum\n "},
	}

	for i, test := range tests {
		output := AlignCenter(test.input, test.width)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleAlignCenter() {
	text := AlignCenter("lorem\nipsum", 9)
	fmt.Println(strings.Replace(text, " ", ".", -1))
	// Output:
	// ..lorem..
	// ..ipsum..
}

func TestCenter(t *testing.T) {
	tests := []struct {
		width    int
		input    string
		expected string
	}{
		{9, "lorem", "  lorem  "},
		{10, "lorem", "  lorem   "},
		{1, "lorem", "lorem"},
		{4, "", "    "},
		{0, "lorem", "lorem"},
	}

	for i, test := range tests {
		output := CenterText(test.input, test.width)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleCenter() {
	fmt.Println("'" + CenterText("lorem", 9) + "'")
	// Output: '  lorem  '
}

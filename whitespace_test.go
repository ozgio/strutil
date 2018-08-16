package strutil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndent(t *testing.T) {
	tests := []struct {
		left     string
		input    string
		expected string
	}{

		{"\t", "Lorem ipsum dolor sit amet", "\tLorem ipsum dolor sit amet"},
		{"\t", "Lorem ipsum\ndolor sit amet\n", "\tLorem ipsum\n\tdolor sit amet\n\t"},
		{"\t", "", ""},
		{"", "Lorem", "Lorem"},
	}

	for i, test := range tests {
		output := Indent(test.input, test.left)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)

	}
}

func ExampleIndent() {
	fmt.Println(Indent("Lorem ipsum\ndolor sit amet", " > "))
	// Output:
	//  > Lorem ipsum
	//  > dolor sit amet
}

func TestGetPadString(t *testing.T) {
	tests := []struct {
		length   int
		input    string
		expected string
	}{

		{0, "-", ""},
		{4, "", ""},
		{4, "-", "----"},
		{4, ".-", ".-.-"},
		{4, ".-=", ".-=."},
		{1, ".-", "."},
	}

	for i, test := range tests {
		output := getPadString(test.length, test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		width    int
		input    string
		pad      string
		expected string
	}{
		{10, "lorem", "-", "-----lorem"},
		{5, "lorem", "-", "lorem"},
		{6, "lorem", ".-", ".lorem"},
		{9, "lorem", ".-", ".-.-lorem"},
		{10, "lorem", "", "lorem"},
		{0, "lorem", "-", "lorem"},
		{4, "lorem", "-", "lorem"},
		{4, "", "-", "----"},
		{6, "lorem", ".-=", ".lorem"},
	}

	for i, test := range tests {
		output := PadLeft(test.input, test.width, test.pad)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExamplePadLeft() {
	fmt.Println(PadLeft("lorem", 10, "-"))
	// Output: -----lorem
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		width    int
		input    string
		pad      string
		expected string
	}{
		{10, "lorem", "-", "lorem-----"},
		{5, "lorem", "-", "lorem"},
		{6, "lorem", ".-", "lorem."},
		{9, "lorem", ".-", "lorem.-.-"},
		{10, "lorem", "", "lorem"},
		{0, "lorem", "-", "lorem"},
		{4, "lorem", "-", "lorem"},
		{4, "", "-", "----"},
	}

	for i, test := range tests {
		output := PadRight(test.input, test.width, test.pad)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExamplePadRight() {
	fmt.Println(PadRight("lorem", 10, "-"))
	// Output: lorem-----
}
func TestPad(t *testing.T) {
	tests := []struct {
		width    int
		input    string
		leftPad  string
		rightPad string
		expected string
	}{
		{9, "lorem", "-", "-", "--lorem--"},
		{10, "lorem", ".-", "-.", ".-lorem-.-"},
		{1, "lorem", ".-", "-.", "lorem"},
		{4, "", ".-", "-.", ".--."},
		{10, "lorem", "", "", "lorem"},
		{10, "lorem", "-", "", "-----lorem"},
	}

	for i, test := range tests {
		output := Pad(test.input, test.width, test.leftPad, test.rightPad)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExamplePad() {
	fmt.Println(Pad("lorem", 9, "-", "-"))
	// Output: --lorem--
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
		output := Center(test.input, test.width)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleCenter() {
	fmt.Println("'" + Center("lorem", 9) + "'")
	// Output: '  lorem  '
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
		{"  lorem  \n  ipsum  \n", 10, "     lorem\n     ipsum\n"},
		{"  lorem  \n  ipsum  \n", 1, "lorem\nipsum\n"},
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
		{"", 10, ""},
		{"lorem", 10, "  lorem   "},
		{"lorem\nipsum", 10, "  lorem   \n  ipsum   "},
		{"    lorem", 10, "  lorem   "},
		{"   lorem\n    ipsum", 10, "  lorem   \n  ipsum   "},
		{"  lorem  \n  ipsum  \n", 10, "  lorem   \n  ipsum   \n"},
		{"  lorem  \n  ipsum  \n", 1, "lorem\nipsum\n"},
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

func TestAlign(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		typ      string
		expected string
	}{
		{"  lorem  ", 10, AlignTypeLeft, "lorem  "},
		{"  lorem  ", 10, AlignTypeRight, "     lorem"},
		{"  lorem  ", 10, AlignTypeCenter, "  lorem   "},
		{"  lorem  ", 10, "", "lorem  "},
	}

	for i, test := range tests {
		output := Align(test.input, test.typ, test.width)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleAlign() {
	fmt.Println(Align("  lorem  \n  ipsum  ", AlignTypeRight, 10))
	// Output:
	//      lorem
	//      ipsum
}

func TestExpandTabs(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected string
	}{
		{"", 2, ""},
		{"\t", 0, ""},
		{"\t\n\t\n", 2, "  \n  \n"},
		{"\t\t", 2, "    "},
		{"\tlorem\n\tipsum\n", 2, "  lorem\n  ipsum\n"},
	}

	for i, test := range tests {
		output := ExpandTabs(test.input, test.count)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

package strutil

import (
	"fmt"
	"strings"
	"testing"
)

func TestAlign(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		typ      AlignType
		expected string
	}{
		{"  lorem  ", 10, Left, "lorem  "}, // Existing
		{"  lorem  ", 10, Right, "     lorem"}, // Existing
		{"  lorem  ", 10, Center, "  lorem   "}, // Existing
		{"  lorem  ", 10, "", "  lorem  "}, // Existing (invalid type, effectively default)

		// New test cases for TestAlign
		// Note: Align() first trims the input string.
		{input: "", width: 5, typ: Left, expected: ""}, // Align Left only trims, does not pad to width
		{input: "", width: 5, typ: Right, expected: "     "},
		{input: "", width: 5, typ: Center, expected: "     "},
		{input: "   ", width: 5, typ: Left, expected: ""}, // "   " trims to "", Align Left only trims
		{input: "   ", width: 5, typ: Right, expected: "     "}, // "   " trims to "" then pad
		{input: "   ", width: 5, typ: Center, expected: "     "}, // "   " trims to "" then pad
		{input: "test", width: 0, typ: Left, expected: "test"},
		{input: "test", width: 0, typ: Right, expected: "test"},
		{input: "test", width: 0, typ: Center, expected: "test"},
		{input: "testing", width: 3, typ: Left, expected: "testing"}, // width < len(trimmed_string) returns trimmed_string (AlignLeft only trims)
		{input: "testing", width: 3, typ: Right, expected: "testing"},
		{input: "testing", width: 3, typ: Center, expected: "testing"},
		{input: "  こんにちは世界  ", width: 10, typ: Left, expected: "こんにちは世界  "}, // Align Left only trims leading spaces
		{input: "  こんにちは世界  ", width: 10, typ: Right, expected: "   こんにちは世界"},
		{input: "  こんにちは世界  ", width: 10, typ: Center, expected: " こんにちは世界  "}, // (10-7)/2 = 1.5. Pad left 1, right 2.
		{input: "こんにちは世界", width: 7, typ: Center, expected: "こんにちは世界"},
		{input: "test", width: 6, typ: AlignType("anything_else"), expected: "test"}, // Invalid type returns input string as is (not trimmed based on existing tests)
		{input: "  test  ", width: 6, typ: AlignType("anything_else"), expected: "  test  "}, // Invalid type returns input string as is
	}

	for i, test := range tests {
		output := Align(test.input, test.typ, test.width)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
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
		{"    lorem", "lorem"}, // Existing
		{"   lorem\n    ipsum", "lorem\nipsum"}, // Existing
		{"  lorem  \n  ipsum  \n", "lorem  \nipsum  \n"}, // Existing

		// New test cases for TestAlignLeft
		{input: "", expected: ""},
		{input: "alreadyLeft", expected: "alreadyLeft"},
		{input: "   ", expected: ""}, // Trims to empty
		{input: "  こんにちは世界", expected: "こんにちは世界"},
		{input: "  a\n  b", expected: "a\nb"},
		{input: "\n  a\n  b", expected: "\na\nb"}, // Leading newline
	}

	for i, test := range tests {
		output := AlignLeft(test.input)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
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
		{"    lorem", 10, "     lorem"}, // Existing
		{"   lorem\n    ipsum", 10, "     lorem\n     ipsum"}, // Existing
		{"  lorem  \n  ipsum  \n", 10, "     lorem\n     ipsum\n          "}, // Existing
		{"  lorem  \n  ipsum  \n", 1, "lorem\nipsum\n "}, // Existing

		// New test cases for TestAlignRight
		// Note: AlignRight also trims each line before aligning.
		{input: "", width: 5, expected: "     "},
		{input: "   ", width: 2, expected: "  "}, // "   " trims to "", width 2
		{input: "   ", width: 0, expected: ""},   // "   " trims to "", width 0
		{input: "   ", width: 3, expected: "   "}, // "   " trims to "", width 3
		{input: "test", width: 0, expected: "test"}, // width 0 returns trimmed line
		{input: "  test  ", width: 4, expected: "test"}, // width == len(trimmed_line)
		{input: "  こんにちは世界  ", width: 10, expected: "   こんにちは世界"}, // "こんにちは世界" is 7 runes. (10-7)=3 spaces left.
		{input: "  こんにちは世界  ", width: 7, expected: "こんにちは世界"},
		{input: "short\n  longer line", width: 12, expected: "       short\n longer line"},
		{input: "  short\n  longer line  ", width: 5, expected: "short\nlonger line"}, // width < len(trimmed_line)
	}

	for i, test := range tests {
		output := AlignRight(test.input, test.width)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
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
		{"", 10, "          "}, // Existing
		{"lorem", 10, "  lorem   "}, // Existing
		{"lorem\nipsum", 10, "  lorem   \n  ipsum   "}, // Existing
		{"    lorem", 10, "  lorem   "}, // Existing, input is "    lorem", trimmed to "lorem"
		{"   lorem\n    ipsum", 10, "  lorem   \n  ipsum   "}, // Existing
		{"  lorem  \n  ipsum  \n", 10, "  lorem   \n  ipsum   \n          "}, // Existing
		{"  lorem  \n  ipsum  \n", 1, "lorem\nipsum\n "}, // Existing

		// New test cases for TestAlignCenter
		// Note: AlignCenter also trims each line before aligning.
		{input: "   ", width: 2, expected: "  "}, // "   " trims to "", width 2
		{input: "   ", width: 0, expected: ""},   // "   " trims to "", width 0
		{input: "   ", width: 3, expected: "   "}, // "   " trims to "", width 3
		{input: "test", width: 0, expected: "test"}, // width 0 returns trimmed line
		{input: "  test  ", width: 4, expected: "test"}, // width == len(trimmed_line)
		{input: "  こんにちは世界  ", width: 10, expected: " こんにちは世界  "}, // "こんにちは世界" is 7 runes. (10-7)/2 = 1.5. Pad left 1, right 2.
		{input: "  こんにちは世界  ", width: 7, expected: "こんにちは世界"},
		{input: "short\n  longer line", width: 12, expected: "   short    \nlonger line "}, // "short" len 5 -> 3,4 pads. "longer line" len 11 (after trim) -> 0,1 pads
		{input: "  short\n  longer line  ", width: 5, expected: "short\nlonger line"}, // width < len(trimmed_line)
	}

	for i, test := range tests {
		output := AlignCenter(test.input, test.width)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
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
		{9, "lorem", "  lorem  "}, // Existing
		{10, "lorem", "  lorem   "}, // Existing
		{1, "lorem", "lorem"}, // Existing
		{4, "", "    "}, // Existing
		{0, "lorem", "lorem"}, // Existing

		// New test cases for TestCenter (CenterText)
		// Note: CenterText does NOT trim input.
		{width: 5, input: "", expected: "     "},
		{width: 5, input: "   ", expected: "     "}, // input "   " (len 3), width 5. rem=2. left=1, right=1. -> "     "
		{width: 2, input: "   ", expected: "   "}, // width < len(input) returns input
		{width: 10, input: "こんにちは世界", expected: " こんにちは世界  "}, // "こんにちは世界" is 7 runes. rem=3. left=1, right=2.
		{width: 7, input: "こんにちは世界", expected: "こんにちは世界"}, // width == len(input)
		{width: 7, input: "test", expected: " test  "}, // "test" is 4 runes. rem=3. left=1, right=2.
		{width: 8, input: "test", expected: "  test  "}, // "test" is 4 runes. rem=4. left=2, right=2.
	}

	for i, test := range tests {
		output := CenterText(test.input, test.width)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleCenter() {
	fmt.Println("'" + CenterText("lorem", 9) + "'")
	// Output: '  lorem  '
}

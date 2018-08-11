package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordwrap(t *testing.T) {
	tests := []struct {
		colLen   int
		input    string
		expected string
	}{

		{0, "Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet"},
		{15, "Lorem ipsum\ndolor\nsit amet", "Lorem ipsum\ndolor\nsit amet"},
		{15, "Lorem ipsum dolor sit amet", "Lorem ipsum\ndolor sit amet"},
		{15, "Lorem ipsum, dolor sit amet.", "Lorem ipsum,\ndolor sit amet."},
		{15, "Lorem ipsum, dolor sit amet.\n", "Lorem ipsum,\ndolor sit amet.\n"},
		{15, "\nLorem ipsum, dolor sit amet.\n", "\nLorem ipsum,\ndolor sit amet.\n"},
		{15, "\n   Lorem ipsum, dolor sit amet.\n", "\nLorem ipsum,\ndolor sit amet.\n"},
		{1, "Lorem ipsum, dolor sit amet", "Lorem\nipsum,\ndolor\nsit\namet"},
		{15, "Το Lorem Ipsum είναι απλά ένα κείμενο χωρίς", "Το Lorem Ipsum\nείναι απλά ένα\nκείμενο χωρίς"},
		{15, "", ""},
		{15, "                        ", ""},
		{15, "Lorem ipsum   dolor sit amet", "Lorem ipsum\ndolor sit amet"},
		{15, "Lorem ipsum,   dolor sit amet.", "Lorem ipsum,\ndolor sit amet."},
		{15, "   Lorem ipsum,   dolor sit amet.", "Lorem ipsum,\ndolor sit amet."},
		{15, "Lorem ipsum,dolor sit amet.", "Lorem ipsum,\ndolor sit amet."},
		//{15, "Lorem ipsum,dolor sit amet.   ", "Lorem ipsum,\ndolor sit amet."},
	}

	for i, test := range tests {
		output := Wordwrap(test.input, test.colLen)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}

}

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

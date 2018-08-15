package strutil

import (
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

func TestMapLines(t *testing.T) {
	tests := []struct {
		input    string
		fn       func(string) string
		expected string
	}{
		{"", strings.ToUpper, ""},
		{"\n\n", strings.ToUpper, "\n\n"},
		{"Lorem\nIpsum", strings.ToUpper, "LOREM\nIPSUM"},
	}

	for i, test := range tests {
		output := MapLines(test.input, test.fn)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
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

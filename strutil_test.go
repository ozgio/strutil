package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"abc", "cba"},
		{"a", "a"},
		{"çınar", "ranıç"},
		{"    yağmur", "rumğay    "},
	}

	for i, test := range tests {
		output := Reverse(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func TestReplaceAllToOne(t *testing.T) {
	tests := []struct {
		input    string
		from     []string
		to       string
		expected string
	}{
		{"", []string{"a"}, "b", ""},
		{"lorem", []string{""}, "-", "-l-o-r-e-m-"},
		{"lorem", []string{"lo", "em"}, "", "r"},
		{"a b c a c f", []string{" ", "a", "b"}, "-", "----c---c-f"},
	}

	for i, test := range tests {
		output := ReplaceAllToOne(test.input, test.from, test.to)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

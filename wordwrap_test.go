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

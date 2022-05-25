package strutil

import "testing"

func TestStrAfter(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"Return every word after this word", "word", "after this word"},
		{"it's not contained", "nothing", ""},
	}

	for i, test := range tests {
		output := StrAfter(test.input, test.search)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func TestStrAfterLast(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"Return every word after the word last occurance", "word", "last occurance"},
	}

	for i, test := range tests {
		output := StrAfterLast(test.input, test.search)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

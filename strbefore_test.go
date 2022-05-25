package strutil

import "testing"

func TestStrBefore(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"before this", "this", "before"},
	}

	for i, test := range tests {
		output := StrBefore(test.input, test.search)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func TestStrBeforeLast(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"before this is this", "this", "before this is"},
	}

	for i, test := range tests {
		output := StrBeforeLast(test.input, test.search)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

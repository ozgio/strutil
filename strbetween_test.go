package strutil

import "testing"

func TestStrBetween(t *testing.T) {
	tests := []struct {
		input    string
		a        string
		b        string
		expected string
	}{
		{"between this and that", "this", "that", "and"},
	}

	for i, test := range tests {
		output := StrBetween(test.input, test.a, test.b)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

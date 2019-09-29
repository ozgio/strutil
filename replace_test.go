package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func ExampleReplaceAllToOne() {
	fmt.Println(ReplaceAllToOne("lorem", []string{"lo", "em"}, "x"))
	// Output: xrx
}

func TestSplice(t *testing.T) {
	tests := []struct {
		input     string
		newStr    string
		start     int
		end       int
		mustPanic bool
		expected  string
	}{
		{"lorem", "x", 0, 2, false, "xrem"},
		{"lorem", "", 0, 2, false, "rem"},
		{"", "x", 0, 2, false, ""},
		{"lorem", "x", 4, 5, false, "lorex"},
		{"lorem", "ipsum", 2, 3, false, "loipsumem"},
		{"lorem", "x", 5, 6, true, ""},
		{"lorem", "x", 4, 4, true, ""},
		{"lorem", "x", 4, 3, true, ""},
	}

	for i, test := range tests {
		if test.mustPanic {
			assert.Panicsf(t, func() {
				_ = Splice(test.input, test.newStr, test.start, test.end)
			}, "Test case %d is not successful\n", i)
		} else {
			output := Splice(test.input, test.newStr, test.start, test.end)
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}
}

func ExampleSplice() {
	fmt.Println(Splice("Lorem", "ipsum", 2, 3))
	// Output: Loipsumem
}

package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	tests := []struct {
		input     string
		start     int
		end       int
		expected  string
		mustPanic bool
	}{
		{"lorem", 0, 1, "l", false},
		{"", 0, 1, "", true},
		{"lorem", 0, 5, "lorem", false},
		{"lorem", 0, 10, "", true},
		{"lorem", -1, 4, "", true},
		{"lorem", 9, 10, "", true},
		{"lorem", 4, 3, "", true},
		{"Υπάρχουν", 1, 4, "πάρ", false},
		{"Υπάρχουν", 1, 0, "πάρχουν", false},
		{"Υπάρχουν", 1, 9, "", true},
		{"žůžo", 1, 4, "ůžo", false},
	}

	for i, test := range tests {
		if test.mustPanic {
			assert.Panicsf(t, func() {
				_ = Substring(test.input, test.start, test.end)
			}, "Test case %d is not successful\n", i)
		} else {
			output := Substring(test.input, test.start, test.end)
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}

}

func ExampleSubstring() {
	fmt.Println(Substring("Υπάρχουν", 1, 4))
	// Output: πάρ
}

func ExampleSubstring_tillTheEnd() {
	fmt.Println(Substring("Υπάρχουν", 1, 0))
	// Output: πάρχουν
}

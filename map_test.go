package strutil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func ExampleMapLines() {
	fmt.Println(MapLines("Lorem\nIpsum", strings.ToUpper))
	// Output:
	// LOREM
	// IPSUM
}

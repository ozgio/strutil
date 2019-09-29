package strutil

import (
	"fmt"
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
		{"\t", "", "\t"},
		{"", "Lorem", "Lorem"},
	}

	for i, test := range tests {
		output := Indent(test.input, test.left)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)

	}
}

func ExampleIndent() {
	fmt.Println(Indent("Lorem ipsum\ndolor sit amet", " > "))
	// Output:
	//  > Lorem ipsum
	//  > dolor sit amet
}

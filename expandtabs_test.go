package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func ExampleExpandTabs() {
	fmt.Printf("%s", ExpandTabs("\tlorem\n\tipsum", 2))
	//Output:
	//   lorem
	//   ipsum
}

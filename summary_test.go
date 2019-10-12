package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummary(t *testing.T) {
	tests := []struct {
		colLen   int
		end      string
		input    string
		expected string
	}{
		{15, "...", "", ""},
		{0, "...", "Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet"},
		{12, "...", "Lorem ipsum dolor sit amet", "Lorem ipsum..."},
		{12, "...", "Lorem\nipsum dolor sit amet", "Lorem..."},
		{12, "...", "Lorem\tipsum\tdolor sit amet", "Lorem\tipsum..."},
		{10, "...", "Lorem ipsum dolor sit amet", "Lorem..."},
		{10, "...", "Lorem\nipsum dolor sit amet", "Lorem..."},
		{15, "...", "Lorem ipsum", "Lorem ipsum"},
		{5, "...", "Lorem ipsum", "Lorem..."},
		{4, "...", "Lorem ipsum", "Lore..."},
		{15, "...", "Lorem         ipsum", "Lorem..."},
	}

	for i, test := range tests {
		output := Summary(test.input, test.colLen, test.end)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleSummary() {
	fmt.Println(Summary("Lorem ipsum dolor sit amet.", 12, "..."))
	// Output: Lorem ipsum...
}

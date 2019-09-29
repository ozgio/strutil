package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTile(t *testing.T) {
	tests := []struct {
		pattern  string
		length   int
		expected string
	}{
		{"", 10, ""},
		{"-", 10, "----------"},
		{"-.", -1, ""},
		{"-.", 0, ""},
		{"-.", 2, "-."},
		{"-.", 1, "-"},
		{"-.", 3, "-.-"},
		{"-.", 4, "-.-."},
		{"-৹", 0, ""},
		{"-৹", 2, "-৹"},
		{"-৹", 1, "-"},
		{"-৹", 3, "-৹-"},
		{"-৹", 4, "-৹-৹"},
	}

	for i, test := range tests {
		output := Tile(test.pattern, test.length)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

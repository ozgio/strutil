package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordwrap(t *testing.T) {
	tests := []struct {
		colLen         int
		breakLongWords bool
		input          string
		expected       string
	}{

		{0, false, "Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet"},
		{15, false, "Lorem ipsum\ndolor\nsit amet", "Lorem ipsum\ndolor\nsit amet"},
		{15, false, "Lorem ipsum dolor sit amet", "Lorem ipsum\ndolor sit amet"},
		{15, false, "Lorem ipsum     \n   dolor sit amet\n    consectetur", "Lorem ipsum    \n\n   dolor sit\namet\n    consectetur"},
		{15, false, "Lorem ipsum, dolor sit amet.", "Lorem ipsum,\ndolor sit amet."},
		{15, false, "Lorem ipsum, dolor sit amet.\n", "Lorem ipsum,\ndolor sit amet.\n"},
		{15, false, "\nLorem ipsum, dolor sit amet.\n", "\nLorem ipsum,\ndolor sit amet.\n"},
		{15, false, "\n   Lorem ipsum, dolor sit amet.\n", "\n   Lorem ipsum,\ndolor sit amet.\n"},
		{1, false, "Lorem ipsum, dolor sit amet", "Lorem\nipsum,\ndolor\nsit\namet"},
		{1, true, "Lorem ipsum, dolor sit amet", "L\no\nr\ne\nm\ni\np\ns\nu\nm\n,\nd\no\nl\no\nr\ns\ni\nt\na\nm\ne\nt"},
		{6, true, "Loremipsum, dolorsitamet", "Loremi\npsum,\ndolors\nitamet"},
		{3, true, "Lorem ipsum, dolor sit amet", "Lor\nem\nips\num,\ndol\nor\nsit\name\nt"},
		{15, false, "Το Lorem Ipsum είναι απλά ένα κείμενο χωρίς", "Το Lorem Ipsum\nείναι απλά ένα\nκείμενο χωρίς"},
		{15, false, "", ""},
		{15, false, "                        ", "               \n        "},
		{15, false, "Lorem ipsum   dolor sit amet", "Lorem ipsum  \ndolor sit amet"},
		{15, false, "Lorem ipsum,   dolor sit amet.", "Lorem ipsum,  \ndolor sit amet."},
		{15, false, "   Lorem ipsum,   dolor sit amet.", "   Lorem ipsum,\n  dolor sit\namet."},
		{15, false, "Lorem ipsum,dolor sit amet.", "Lorem\nipsum,dolor sit\namet."},
		{15, false, "Lorem ipsum,dolor sit amet.   ", "Lorem\nipsum,dolor sit\namet.   "},
		{7, true, "Lorem ipsum,dolor sit amet.   ", "Lorem\nipsum,d\nolor\nsit\namet.  "},
	}

	for i, test := range tests {
		output := WordWrap(test.input, test.colLen, test.breakLongWords)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}

}

func ExampleWordWrap() {
	fmt.Println(WordWrap("Lorem ipsum, dolor sit amet.", 15, false))
	// Output:
	// Lorem ipsum,
	// dolor sit amet.

}

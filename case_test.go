package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"lorem", "lorem"},
		{"lorem ipsum", "lorem_ipsum"},
		{"Lorem Ipsum", "lorem_ipsum"},
		{"", ""},
		{" ", ""},
	}

	for i, test := range tests {
		output := ToSnakeCase(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleToSnakeCase() {
	fmt.Println(ToSnakeCase("Lorem Ipsum"))
	// Output: lorem_ipsum
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"lorem", "lorem"},
		{"lorem ipsum", "loremIpsum"},
		{"Lorem Ipsum", "LoremIpsum"},
		{"bay ğ", "bayĞ"},
		{"", ""},
		{" ", ""},
	}

	for i, test := range tests {
		output := ToCamelCase(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleToCamelCase() {
	fmt.Println(ToCamelCase("long live motörhead"))
	//Output: longLiveMotörhead
}

func TestSplitCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"lorem", []string{"lorem"}},
		{"loremIpsum", []string{"lorem", "Ipsum"}},
		{"binaryJSONAbstractWriter", []string{"binary", "JSON", "Abstract", "Writer"}},
		{"bayĞe", []string{"bay", "Ğe"}},
		{"", []string{""}},
		{" ", []string{""}},
		{"HTML", []string{"HTML"}},
		{"AClass", []string{"A", "Class"}},
		{"year2000", []string{"year", "2000"}},
		{"year2000OfWorld", []string{"year", "2000", "Of", "World"}},
		{"year2000s", []string{"year", "2000s"}},
		{"yearX2000s", []string{"year", "X", "2000s"}},
	}

	for i, test := range tests {
		output := SplitCamelCase(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleSplitCamelCase() {
	fmt.Printf("%#v\n", SplitCamelCase("binaryJSONAbstractWriter"))
	// Output: []string{"binary", "JSON", "Abstract", "Writer"}
}

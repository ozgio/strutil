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
	fmt.Println(SplitCamelCase("binaryJSONAbstractWriter"))
	// Output: [binary JSON Abstract Writer]
}

func TestSlugifySpecial(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"  ", ""},
		{"lorem ", "lorem"},
		{"We löve Motörhead", "we-love-motorhead"},
		{"žůžo", "zuzo"},
		{"yağmur    --    yağarken", "yagmur-yagarken"},
		{"résumé", "resume"},
		{"Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero, written in 45 BC. ", "lorem-ipsum-comes-from-sections-1-10-32-and-1-10-33-of-de-finibus-bonorum-et-malorum-the-extremes-of-good-and-evil-by-cicero-written-in-45-bc"},
	}

	for i, test := range tests {
		output := SlugifySpecial(test.input, "-")
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleSlugifySpecial() {
	fmt.Println(SlugifySpecial("We löve Motörhead", "_"))
	// Output: we_love_motorhead
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"  ", ""},
		{"lorem ", "lorem"},
		{"We löve Motörhead", "we-love-motorhead"},
		{"žůžo", "zuzo"},
		{"yağmur    --    yağarken", "yagmur-yagarken"},
		{"résumé", "resume"},
		{"Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero, written in 45 BC. ", "lorem-ipsum-comes-from-sections-1-10-32-and-1-10-33-of-de-finibus-bonorum-et-malorum-the-extremes-of-good-and-evil-by-cicero-written-in-45-bc"},
	}

	for i, test := range tests {
		output := Slugify(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleSlugify() {
	fmt.Println(Slugify("We löve Motörhead"))
	// Output: we-love-motorhead
}

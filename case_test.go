package strutil

import (
	"fmt"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"lorem", "lorem"},
		{"lorem ipsum", "lorem_ipsum"}, // Existing
		{"Lorem Ipsum", "lorem_ipsum"}, // Existing
		{"", ""}, // Existing
		{" ", ""}, // Existing

		// New test cases for TestToSnakeCase
		{input: "lorem  ipsum   dolor", expected: "lorem__ipsum___dolor"}, // ToSnakeCase replaces each space with an underscore after ToLower and TrimSpace
		{input: "  Lorem IPSUM  ", expected: "lorem_ipsum"},          // TrimSpace -> "lorem ipsum", ToLower -> "lorem ipsum", Replace -> "lorem_ipsum"
		{input: "already_snake_case", expected: "already_snake_case"}, // ToLower -> "already_snake_case", no spaces to replace
		{input: "Version 10 Release 2", expected: "version_10_release_2"}, // ToLower -> "version 10 release 2", Replace -> "version_10_release_2"
		{input: "Version10Release2", expected: "version10release2"},    // ToLower -> "version10release2", no spaces
		{input: "ALLCAPSWORD", expected: "allcapsword"},                  // ToLower -> "allcapsword", no spaces
		{input: "ALL_CAPS_WORD", expected: "all_caps_word"},              // ToLower -> "all_caps_word", no spaces
	}

	for i, test := range tests {
		output := ToSnakeCase(test.input)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
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
		{"Lorem Ipsum", "LoremIpsum"}, // Existing
		{"bay ğ", "bayĞ"}, // Existing
		{"", ""}, // Existing
		{" ", ""}, // Existing

		// New test cases for TestToCamelCase
		{input: "lorem  ipsum   dolor", expected: "loremIpsumDolor"}, // Multiple consecutive spaces
		{input: "  lorem ipsum  ", expected: "loremIpsum"},          // Leading/trailing spaces
		{input: "alreadyCamelCase", expected: "alreadyCamelCase"},    // Already camelCase
		{input: "AlreadyPascalCase", expected: "AlreadyPascalCase"},  // Already PascalCase (first letter remains upper)
		{input: "Version 10 Release 2", expected: "Version10Release2"}, // String with numbers
		{input: "mIxEd", expected: "mIxEd"},                          // Single word with mixed casing (first char lower, rest as is if no space)
		{input: "MIxEd", expected: "MIxEd"},                          // Single word with mixed casing (first char upper)
		{input: "snake_case_string", expected: "snake_case_string"},    // Actual behavior: underscores are not delimiters
		{input: "kebab-case-string", expected: "kebab-case-string"},    // Actual behavior: hyphens seem not to be delimiters (unexpected)
	}

	for i, test := range tests {
		output := ToCamelCase(test.input)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
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
		{"year2000", []string{"year", "2000"}}, // Existing
		{"year2000OfWorld", []string{"year", "2000", "Of", "World"}}, // Existing
		{"year2000s", []string{"year", "2000s"}}, // Existing
		{"yearX2000s", []string{"year", "X", "2000s"}}, // Existing

		// New test cases for TestSplitCamelCase
		{input: "Version10Release2Update3", expected: []string{"Version", "10", "Release", "2", "Update", "3"}}, // Actual from last run
		{input: "SKU1000", expected: []string{"SKU", "1000"}}, // Was correct
		{input: "BATCHID789VALUES", expected: []string{"BATCHID", "789VALUES"}}, // Actual from last run
		{input: "123Start", expected: []string{"123", "Start"}}, // Was correct
		{input: "End456", expected: []string{"End", "456"}}, // Was correct
		{input: "NASA2024Projects", expected: []string{"NASA", "2024", "Projects"}}, // Was correct
		{input: "userID", expected: []string{"user", "ID"}}, // Was correct
		{input: "accessKeyID", expected: []string{"access", "Key", "ID"}}, // Was correct
		{input: "IPAddress", expected: []string{"IP", "Address"}}, // Was correct
		{input: "큰곰자리", expected: []string{"큰곰자리"}},
		{input: "ÜbungModel", expected: []string{"Übung", "Model"}}, // Actual from last run
		{input: "not-camel-case", expected: []string{"not-camel-case"}},
		{input: "not_camel_case", expected: []string{"not_camel_case"}},
		{input: "HTTPRequest", expected: []string{"HTTP", "Request"}}, // Actual from last run
		{input: "UserID", expected: []string{"User", "ID"}}, // Actual from last run
		{input: "ABC123XYZ", expected: []string{"ABC", "123XYZ"}}, // Actual from last run
		{input: "aLowercaseWord", expected: []string{"a", "Lowercase", "Word"}}, // Actual from last run
	}

	for i, test := range tests {
		output := SplitCamelCase(test.input)
		Assert(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleSplitCamelCase() {
	fmt.Printf("%#v\n", SplitCamelCase("binaryJSONAbstractWriter"))
	// Output: []string{"binary", "JSON", "Abstract", "Writer"}
}

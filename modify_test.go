package strutil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"abc", "cba"},
		{"a", "a"},
		{"çınar", "ranıç"},
		{"    yağmur", "rumğay    "},
		{"επαγγελματίες", "ςείταμλεγγαπε"},
	}

	for i, test := range tests {
		output := Reverse(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("επαγγελματίες"))
	// Outputs: ςείταμλεγγαπε
}

func TestReplaceAllToOne(t *testing.T) {
	tests := []struct {
		input    string
		from     []string
		to       string
		expected string
	}{
		{"", []string{"a"}, "b", ""},
		{"lorem", []string{""}, "-", "-l-o-r-e-m-"},
		{"lorem", []string{"lo", "em"}, "", "r"},
		{"a b c a c f", []string{" ", "a", "b"}, "-", "----c---c-f"},
	}

	for i, test := range tests {
		output := ReplaceAllToOne(test.input, test.from, test.to)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleReplaceAllToOne() {
	fmt.Println(ReplaceAllToOne("lorem", []string{"lo", "em"}, "x"))
	// Outputs: xrx
}

func TestRemoveAccents(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"lorem", "lorem"},
		{"motörhead", "motorhead"},
		{"žůžo", "zuzo"},
		{"yağmur", "yagmur"},
		{"résumé", "resume"},
		{"çınar", "cinar"},
		{"ÄØääkkönen", "AOaakkonen"},
		{"£ $ ä", "£ $ a"},
		{"ßąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüűûñÿýçżźž", "ssaaaaaaaaaccceeeeeghiiiijllnnoooooooossssstttuuuuuunyyczzz"}, //taken from github.com/epeli/underscore.string
	}

	for i, test := range tests {
		output, _, _ := RemoveAccents(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleRemoveAccents() {
	fmt.Println(RemoveAccents("ßąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüűûñÿýçżźž"))
	// Outputs: ssaaaaaaaaaccceeeeeghiiiijllnnoooooooossssstttuuuuuunyyczzz
}

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
	// Outputs:
	// LOREM
	// IPSUM
}

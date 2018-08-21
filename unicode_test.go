package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	output, _, _ := RemoveAccents("ßąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüűûñÿýçżźž")
	fmt.Println(output)
	// Output: ssaaaaaaaaaccceeeeeghiiiijllnnoooooooossssstttuuuuuunyyczzz
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

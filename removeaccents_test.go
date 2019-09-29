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
		{"ÄØääkkönen", "AOEaakkonen"},
		{"£ $ ä", "£ $ a"},
		{"ßąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüűûñÿýçżźž", "ssaaaaaaaaeaccceeeeeghiiiijllnnoooooodjoessssstttuuuuuunyyczzz"},
	}

	for i, test := range tests {
		output := RemoveAccents(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleRemoveAccents() {
	output := RemoveAccents("ßąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüűûñÿýçżźž")
	fmt.Println(output)
	// Output: ssaaaaaaaaeaccceeeeeghiiiijllnnoooooodjoessssstttuuuuuunyyczzz
}

const removeAccentsTestStringAll = "ßąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüóöőôõðøçżźž"
const removeAccentsTestStringBalanced = "ssaßą àáa aaaaaaeacccëêĝhiiiijl lnnoooóö sssuuuuunyyczzuuuukðø"
const removeAccentsTestStringNone = "ssaaaaaaaaeaccceeeeeghiiiijllnnoooooodjoessssstttuuuuuunyyczzz"

func BenchmarkRemoveAccentsWorse(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = RemoveAccents(removeAccentsTestStringAll)
	}
	_ = s
}

func BenchmarkRemoveAccentsBalanced(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = RemoveAccents(removeAccentsTestStringBalanced)
	}
	_ = s
}
func BenchmarkRemoveAccentsBest(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = RemoveAccents(removeAccentsTestStringNone)
	}
	_ = s
}

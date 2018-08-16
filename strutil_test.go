package strutil

import (
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
	}

	for i, test := range tests {
		output := Reverse(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
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

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"    ", 0},
		{"he/she/it", 1}, //FIXME should be 3
		{"...,", 1},      //FIXME should be 3
		{`It is not known exactly when the text obtained its current standard form; it may have been as late as the 1960s. Dr. Richard McClintock, a Latin scholar who was the publications director at Hampden–Sydney College in Virginia, discovered the source of the passage sometime before 1982 while searching for instances of the Latin word "consectetur" ("that [he/she/it] pursue", subjunctive), rarely used in classical literature.[2][a] The physical source of the lorem ipsum text may be the 1914 Loeb Classical Library Edition of the De Finibus, where the Latin text, presented on the left-hand (even) pages, breaks off on page 34 with "Neque porro quisquam est qui do-" and continues on page 36 with "lorem ipsum ...", suggesting that the galley type of that page was mixed up to make the dummy text seen today.`, 134}, //FIXME should be 135
		{`Υπάρχουν αρκετές παραλλαγές, μερικές από τις οποίες δεν έχουν καμία σχέση με το "κλασικό" κείμενο. `, 15},
	}

	for i, test := range tests {
		output := CountWords(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func TestSubstring(t *testing.T) {
	tests := []struct {
		input     string
		start     int
		end       int
		expected  string
		mustPanic bool
	}{
		{"lorem", 0, 1, "l", false},
		{"", 0, 1, "", true},
		{"lorem", 0, 5, "lorem", false},
		{"lorem", 0, 10, "", true},
		{"Υπάρχουν", 1, 4, "πάρ", false},
		{"žůžo", 1, 4, "ůžo", false},
	}

	for i, test := range tests {
		if test.mustPanic {
			assert.Panicsf(t, func() {
				_ = Substring(test.input, test.start, test.end)
			}, "Test case %d is not successful\n", i)
		} else {
			output := Substring(test.input, test.start, test.end)
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}

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

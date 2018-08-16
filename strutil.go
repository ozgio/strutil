//Package strutil provides string utilities for utf8 encoded strings. It is
//complemantary to builtin strings package.
package strutil

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// CountWords count the words approximately based on strings.Fields function.
func CountWords(str string) int {
	var count int
	arr := strings.Fields(str)
	count = len(arr)

	return count
}

// Substring gets a part of the string between start and end. If end is -1,
// end is taken as the length of the string.
//
// It is UTF8 safe version of using slice notations in strings. It panics
// when the indexes are out of range. String length can be get with
// UTF8Len function before using Substring
func Substring(str string, start int, end int) string {
	runes := []rune(str)
	if start < 0 || start > len(runes)-1 {
		panic(fmt.Sprintf("start (%d) is out of range (%d)", start, len(runes)))
	}
	if end == -1 {
		return string(runes[start:])
	}

	if end <= start || end > len(runes) {
		panic(fmt.Sprintf("end (%d) is out of range (%d)", end, len(runes)))

	}
	return string(runes[start:end])
}

// UTF8Len is an alias of utf8.RuneCountInString which returns the number of
// runes in s. Erroneous and short encodings are treated as single runes of
// width 1 byte.
var UTF8Len = utf8.RuneCountInString

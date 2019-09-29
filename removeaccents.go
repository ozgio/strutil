package strutil

import (
	"strings"
)

// RemoveAccents removes accents from the letters. The resuting string only has
// the letters from English alphabet.
// It may not be work as expected for some specific letters. Please create an
// issue for these situations.
func RemoveAccents(str string) string {
	var buff strings.Builder
	buff.Grow(len(str))
	for _, r := range str {
		buff.WriteString(normalizeRune(r))
	}

	return buff.String()
}

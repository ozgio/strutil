package strutil

import (
	"strings"
)

// RemoveAccents removes accents from the string. The resulting string only has
// the letters from English alphabet. For example, "résumé" becomes "resume".
// It may not work as expected for some specific letters. Please create an
// issue for these situations.
func RemoveAccents(str string) string {
	var buff strings.Builder
	buff.Grow(len(str))
	for _, r := range str {
		buff.WriteString(normalizeRune(r))
	}

	return buff.String()
}

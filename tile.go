package strutil

import "strings"

// Tile writes the pattern until the length reaches the 'length'
// It returns empty string if the pattern is "" or length <= 0
func Tile(pattern string, length int) string {
	patLen := Len(pattern)
	if len(pattern) == 0 || length <= 0 {
		return ""
	}

	var buff strings.Builder
	for i := 0; i < length; i += patLen {
		buff.WriteString(pattern)
	}
	return Substring(buff.String(), 0, length)

}

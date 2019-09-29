package strutil

import "unicode/utf8"

// IsASCII checks if all the characters in string are in standard ASCII table
// It is taken from strings.Fields function
func IsASCII(s string) bool {
	// setBits is used to track which bits are set in the bytes of s.
	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
		setBits |= s[i]
	}

	return setBits < utf8.RuneSelf
}

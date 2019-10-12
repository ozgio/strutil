package strutil

import (
	"fmt"
)

// MustSubstring gets a part of the string between start and end. If end is 0,
// end is taken as the length of the string.
//
// It is UTF8 safe version of using slice notations in strings. It panics
// when the indexes are out of range. String length can be get with
// UTF8Len function before using Substring. You can use "Substring" if
// you prefer errors to panics.
func MustSubstring(str string, start int, end int) string {
	res, err := Substring(str, start, end)
	if err != nil {
		panic(err)
	}
	return res
}

// Substring gets a part of the string between start and end. If end is 0,
// end is taken as the length of the string.
//
// MustSubstring can be used for the cases where the boundaries are wwll known and/or panics are
// acceptable
//
// It is UTF8 safe version of using slice notations in strings.
func Substring(str string, start int, end int) (string, error) {
	if start < 0 || start >= len(str) {
		return "", fmt.Errorf("start (%d) is out of range", start)
	}
	if end != 0 && end <= start {
		return "", fmt.Errorf("end (%d) cannot be equal to or smaller than start (%d)", end, start)
	}
	if end > len(str) {
		return "", fmt.Errorf("end (%d) is out of range", end)
	}

	var startByte = -1
	var runeIndex int
	for i := range str {
		if runeIndex == start {
			startByte = i
			if end == 0 {
				return str[startByte:], nil
			}
		}
		if end != 0 && runeIndex == end {
			return str[startByte:i], nil
		}
		runeIndex++
	}

	if startByte < 0 {
		return "", fmt.Errorf("start (%d) is out of range (%d)", start, runeIndex)
	}

	if end == runeIndex {
		return str[startByte:], nil
	}

	return "", fmt.Errorf("end (%d) is out of range (%d)", end, runeIndex)
}

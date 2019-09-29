package strutil

import "strings"

func getPadString(length int, pad string) string {
	if len(pad) == 0 || length <= 0 {
		return ""
	}
	allPad := strings.Repeat(pad, length/Len(pad))
	remaining := length - Len(allPad)
	return allPad + pad[:remaining]
}

// PadLeft left pads a string str with "pad". The string is padded to
// the size of width.
func PadLeft(str string, width int, pad string) string {
	return getPadString(width-Len(str), pad) + str
}

// PadRight right pads a string str with "pad". The string is padded to
// the size of width.
func PadRight(str string, width int, pad string) string {
	return str + getPadString(width-Len(str), pad)
}

// Pad left and right pads a string str with leftPad and rightPad. The string
// is padded to the size of width.
func Pad(str string, width int, leftPad string, rightPad string) string {
	switch {
	case Len(leftPad) == 0:
		return PadRight(str, width, rightPad)
	case Len(rightPad) == 0:
		return PadLeft(str, width, leftPad)
	}
	padLen := (width - Len(str)) / 2
	return getPadString(padLen, leftPad) + str + getPadString(width-Len(str)-padLen, rightPad)
}

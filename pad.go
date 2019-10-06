package strutil

// PadLeft left pads a string str with "pad". The string is padded to
// the size of width.
func PadLeft(str string, width int, pad string) string {
	return Tile(pad, width-Len(str)) + str
}

// PadRight right pads a string str with "pad". The string is padded to
// the size of width.
func PadRight(str string, width int, pad string) string {
	return str + Tile(pad, width-Len(str))
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
	return Tile(leftPad, padLen) + str + Tile(rightPad, width-Len(str)-padLen)
}

package strutil

import (
	"strings"
)

// Indent indents every line of string str with the left parameter
// For empty strings it returns "". Empty lines are indented too.
func Indent(str string, left string) string {
	if str == "" {
		return str
	}
	return left + strings.Replace(str, "\n", "\n"+left, -1)
}

func getPadString(length int, pad string) string {
	if len(pad) == 0 || length <= 0 {
		return ""
	}
	allPad := strings.Repeat(pad, length/UTF8Len(pad))
	remaining := length - UTF8Len(allPad)
	return allPad + pad[:remaining]
}

// PadLeft left pads a string str with "pad". The string is padded to
// the size of width.
func PadLeft(str string, width int, pad string) string {
	return getPadString(width-UTF8Len(str), pad) + str
}

// PadRight right pads a string str with "pad". The string is padded to
// the size of width.
func PadRight(str string, width int, pad string) string {
	return str + getPadString(width-UTF8Len(str), pad)
}

// Pad left and right pads a string str with leftPad and rightPad. The string
// is padded to the size of width.
func Pad(str string, width int, leftPad string, rightPad string) string {
	switch {
	case UTF8Len(leftPad) == 0:
		return PadRight(str, width, rightPad)
	case UTF8Len(rightPad) == 0:
		return PadLeft(str, width, leftPad)
	}
	padLen := (width - UTF8Len(str)) / 2
	return getPadString(padLen, leftPad) + str + getPadString(width-UTF8Len(str)-padLen, rightPad)
}

// Center centers the text by adding spaces to the left and right.
func Center(str string, width int) string {
	return Pad(str, width, " ", " ")
}

// AlignLeft aligns str to the left. It actually trims and right pads all the lines
// in the text with space to the size of width.
func AlignLeft(str string, width int) string {
	return MapLines(str, func(line string) string {
		line = strings.TrimLeft(line, " ")
		return PadRight(line, width, " ")
	})
}

// AlignRight aligns str to the right. It actually trims and left pads all the lines
// in the text with space to the size of width.
func AlignRight(str string, width int) string {
	return MapLines(str, func(line string) string {
		line = strings.Trim(line, " ")
		return PadLeft(line, width, " ")
	})
}

// AlignCenter centers str. It trims and then centers all the lines in the
// text with space
func AlignCenter(str string, width int) string {
	return MapLines(str, func(line string) string {
		line = strings.Trim(line, " ")
		return Center(line, width)
	})
}

// Align type to use with align function
const (
	AlignTypeCenter = "center"
	AlignTypeLeft   = "left"
	AlignTypeRight  = "right"
)

// Align aligns string to the "typ" which should be one of
//  - strutil.AlignTypeCenter
//  - strutil.AlignTypeLeft
//  - strutil.AlignTypeRight
func Align(str string, typ string, width int) string {
	switch typ {
	case AlignTypeCenter:
		return AlignCenter(str, width)
	case AlignTypeRight:
		return AlignRight(str, width)
	default:
		return AlignLeft(str, width)
	}
}

//ExpandTabs convert tabs the spaces with the length of count
func ExpandTabs(str string, count int) string {
	return strings.Replace(str, "\t", strings.Repeat(" ", count), -1)
}

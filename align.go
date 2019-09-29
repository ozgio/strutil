package strutil

import (
	"strings"
)

type AlignType string

// Align type to use with align function
const (
	Center AlignType = "center"
	Left   AlignType = "left"
	Right  AlignType = "right"
)

// Align aligns string to the "typ" which should be one of
//  - strutil.Center
//  - strutil.Left
//  - strutil.Right
func Align(str string, typ AlignType, width int) string {
	switch typ {
	case Center:
		return AlignCenter(str, width)
	case Right:
		return AlignRight(str, width)
	default:
		return AlignLeft(str, width)
	}
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
		return CenterText(line, width)
	})
}

// CenterText centers the text by adding spaces to the left and right.
func CenterText(str string, width int) string {
	return Pad(str, width, " ", " ")
}

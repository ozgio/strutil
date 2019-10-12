package strutil

import (
	"strings"
)

//AlignType text align variable like center or left
type AlignType string

// Align type to use with align function
const (
	Center AlignType = "center"
	Left   AlignType = "left"
	Right  AlignType = "right"
)

// Align aligns string to the "alignTo" which should be one of
// - strutil.Center
// - strutil.Left
// - strutil.Right
func Align(str string, alignTo AlignType, width int) string {
	switch alignTo {
	case Center:
		return AlignCenter(str, width)
	case Right:
		return AlignRight(str, width)
	case Left:
		return AlignLeft(str)
	default:
		return str
	}
}

// AlignLeft aligns string to the left. To achieve that it left trims every line.
func AlignLeft(str string) string {
	return MapLines(str, func(line string) string {
		return strings.TrimLeft(line, " ")
	})
}

// AlignRight aligns string to the right. It trims and left pads all the lines
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
// It assumes the text is one line. For multiple lines use AlignCenter.
func CenterText(str string, width int) string {
	return Pad(str, width, " ", " ")
}

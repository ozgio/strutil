package strutil

import (
	"errors"
	"strings"
)

// Box9Slice is used by Box functions to draw frames around text content by
// defining the corner and edge characters. See DefaultBox9Slice for an
// example
type Box9Slice struct {
	Top         string
	TopRight    string
	Right       string
	BottomRight string
	Bottom      string
	BottomLeft  string
	Left        string
	TopLeft     string
}

// DefaultBox9Slice defines the character object to use with "CustomBox".
// It is used as Box9Slice object in "Box" function.
//
// Usage:
// CustomBox("Hello World", 20, AligntTypeCenter, DefaultBox9Slice)
//
// Outputs:
//   ┌──────────────────┐
//   │   Hello World    │
//   └──────────────────┘
// </code>
var DefaultBox9Slice = Box9Slice{
	Top:         "─",
	TopRight:    "┐",
	Right:       "│",
	BottomRight: "┘",
	Bottom:      "─",
	BottomLeft:  "└",
	Left:        "│",
	TopLeft:     "┌",
}

// SimpleBox9Slice defines a character set to use with CustomBox. It uses
// only simple ASCII chaaracters
//
// Usage:
// CustomBox("Hello World", 20, AligntTypeCenter, SimpleBox9Slice)
//
// Outputs:
//   +------------------+
//   |   Hello World    |
//   +------------------+
var SimpleBox9Slice = Box9Slice{
	Top:         "-",
	TopRight:    "+",
	Right:       "|",
	BottomRight: "+",
	Bottom:      "-",
	BottomLeft:  "+",
	Left:        "|",
	TopLeft:     "+",
}

// CustomBox creates a frame with "content" in it. Characters in the frame is specified by "chars".
// "align" sets the alignment of the content. It must be one of the strutil.AlignType* constants.
// There are 2 premade Box9Slice objects: strutil.DefaultBox9Slice or strutil.SimpleBox9Slice.
// CustomBox wrap the lines with strutil.WordWrap before placing it.
//
// Usage:
// CustomBox("Hello World", 20, AligntTypeCenter, SimpleBox9Slice)
func CustomBox(content string, width int, align string, chars Box9Slice) (string, error) {
	var buff strings.Builder

	var topInsideWidth = width - UTF8Len(chars.TopLeft) - UTF8Len(chars.TopRight)
	var middleInsideWidth = width - UTF8Len(chars.Left) - UTF8Len(chars.Right)
	var bottomInsideWidth = width - UTF8Len(chars.BottomLeft) - UTF8Len(chars.BottomRight)
	if topInsideWidth < 1 || middleInsideWidth < 1 || bottomInsideWidth < 1 {
		return "", errors.New("no enough width")
	}

	content = Wordwrap(content, middleInsideWidth, true)

	buff.WriteString(chars.TopLeft)
	buff.WriteString(strings.Repeat(chars.Top, topInsideWidth))
	buff.WriteString(chars.TopRight)
	buff.WriteString("\n")

	buff.WriteString(MapLines(content, func(line string) string {
		line = Align(line, align, middleInsideWidth)
		if align == AlignTypeLeft {
			line = PadRight(line, middleInsideWidth, " ")
		}
		if line == "" {
			line = strings.Repeat(" ", middleInsideWidth)
		}
		return chars.Left + line + chars.Right
	}))
	buff.WriteString("\n")

	buff.WriteString(chars.BottomLeft)
	buff.WriteString(strings.Repeat(chars.Bottom, bottomInsideWidth))
	buff.WriteString(chars.BottomRight)

	return buff.String(), nil
}

// Box creates a frame with "content" in it. DefaultBox9Slice object is used to
// define characters in the frame. "align" sets the alignment of the content.
// It must be one of the strutil.AlignType* constants.
//
// Usage:
// Box("Hello World", 20, AligntTypeCenter)
func Box(content string, width int, align string) (string, error) {
	return CustomBox(content, width, align, DefaultBox9Slice)
}

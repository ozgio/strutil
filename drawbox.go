package strutil

import (
	"errors"
	"strings"
)

// Box9Slice is used by DrawBox functions to draw frames around text content by
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

var defaultBox9Slice = Box9Slice{
	Top:         "─",
	TopRight:    "┐",
	Right:       "│",
	BottomRight: "┘",
	Bottom:      "─",
	BottomLeft:  "└",
	Left:        "│",
	TopLeft:     "┌",
}

// DefaultBox9Slice defines the character object to use with "CustomBox".
// It is used as Box9Slice object in "DrawBox" function.
//
// Usage:
// DrawCustomBox("Hello World", 20, AligntTypeCenter, DefaultBox9Slice())
//
// Outputs:
// <code>
//   ┌──────────────────┐
//   │   Hello World    │
//   └──────────────────┘
// </code>
func DefaultBox9Slice() Box9Slice {
	return defaultBox9Slice
}

var simpleBox9Slice = Box9Slice{
	Top:         "-",
	TopRight:    "+",
	Right:       "|",
	BottomRight: "+",
	Bottom:      "-",
	BottomLeft:  "+",
	Left:        "|",
	TopLeft:     "+",
}

// SimpleBox9Slice defines a character set to use with DrawCustomBox. It uses
// only simple ASCII characters
//
// Usage:
//   DrawCustomBox("Hello World", 20, Center, SimpleBox9Slice(), "\n")
//
// Outputs:
//   +------------------+
//   |   Hello World    |
//   +------------------+
func SimpleBox9Slice() Box9Slice {
	return simpleBox9Slice
}

// DrawCustomBox creates a frame with "content" in it. Characters in the frame is specified by "chars".
// "align" sets the alignment of the content. It must be one of the strutil.AlignType constants.
// There are 2 premade Box9Slice objects that can be retrieved by strutil.DefaultBox9Slice() or
// strutil.SimpleBox9Slice()
//
// Usage:
//   DrawCustomBox("Hello World", 20, Center, SimpleBox9Slice(), "\n")
//
// Outputs:
//   +------------------+
//   |   Hello World    |
//   +------------------+
func DrawCustomBox(content string, width int, align AlignType, chars Box9Slice, strNewLine string) (string, error) {
	nl := []byte("\n")
	if strNewLine != "" {
		nl = []byte(strNewLine)
	}

	var topInsideWidth = width - Len(chars.TopLeft) - Len(chars.TopRight)
	var middleInsideWidth = width - Len(chars.Left) - Len(chars.Right)
	var bottomInsideWidth = width - Len(chars.BottomLeft) - Len(chars.BottomRight)
	if topInsideWidth < 1 || middleInsideWidth < 1 || bottomInsideWidth < 1 {
		return "", errors.New("no enough width")
	}

	content = WordWrap(content, middleInsideWidth, true)
	lines := strings.Split(content, "\n")

	var buff strings.Builder
	minNumBytes := (width + 1) * (len(lines) + 2)
	buff.Grow(minNumBytes)

	//top
	buff.WriteString(chars.TopLeft)
	buff.WriteString(Tile(chars.Top, topInsideWidth))
	buff.WriteString(chars.TopRight)
	buff.Write(nl)

	//middle
	left := []byte(chars.Left)
	right := []byte(chars.Right)
	for _, line := range lines {
		line = Align(line, align, middleInsideWidth)
		if align == Left {
			line = PadRight(line, middleInsideWidth, " ")
		}
		if line == "" {
			line = strings.Repeat(" ", middleInsideWidth)
		}

		buff.Write(left)
		buff.WriteString(line)
		buff.Write(right)
		buff.Write(nl)
	}

	//bottom
	buff.WriteString(chars.BottomLeft)
	buff.WriteString(Tile(chars.Bottom, bottomInsideWidth))
	buff.WriteString(chars.BottomRight)

	return buff.String(), nil
}

// DrawBox creates a frame with "content" in it. DefaultBox9Slice object is used to
// define characters in the frame. "align" sets the alignment of the content.
// It must be one of the strutil.AlignType constants.
//
// Usage:
//   DrawBox("Hello World", 20, Center)
//
// Outputs:
//   ┌──────────────────┐
//   │   Hello World    │
//   └──────────────────┘
func DrawBox(content string, width int, align AlignType) (string, error) {
	return DrawCustomBox(content, width, align, defaultBox9Slice, "\n")
}

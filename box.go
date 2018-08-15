package strutil

import (
	"errors"
	"strings"
)

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

var DefaultBox9Slice Box9Slice = Box9Slice{
	Top:         "-",
	TopRight:    "+",
	Right:       "|",
	BottomRight: "+",
	Bottom:      "-",
	BottomLeft:  "+",
	Left:        "|",
	TopLeft:     "+",
}

func CustomBox(content string, width int, align string, chars Box9Slice) (string, error) {
	var buff strings.Builder

	var topInsideWidth = width - len(chars.TopLeft) - len(chars.TopRight)
	var middleInsideWidth = width - len(chars.Left) - len(chars.Right)
	var bottomInsideWidth = width - len(chars.BottomLeft) - len(chars.BottomRight)
	if topInsideWidth < 1 || middleInsideWidth < 1 || bottomInsideWidth < 1 {
		return "", errors.New("there is not enough width")
	}

	buff.WriteString(chars.TopLeft)
	buff.WriteString(strings.Repeat(chars.Top, topInsideWidth))
	buff.WriteString(chars.TopRight)

	buff.WriteString(MapLines(content, func(line string) string {
		return chars.Left + Align(line, align, width) + chars.Right
	}))

	buff.WriteString(chars.BottomLeft)
	buff.WriteString(strings.Repeat(chars.Bottom, bottomInsideWidth))
	buff.WriteString(chars.BottomRight)

	return buff.String(), nil
}

func Box(content string, width int, align string) (string, error) {
	return CustomBox(content, width, align, DefaultBox9Slice)
}

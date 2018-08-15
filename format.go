package strutil

import (
	"strings"
)

//Indent indents every line of string str with the string left
//For empty strings it returns "". Empty lines are indented too.
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
	allPad := strings.Repeat(pad, length/len(pad))
	remaining := length - len(allPad)
	return allPad + pad[:remaining]
}

func PadLeft(str string, width int, pad string) string {
	return getPadString(width-len(str), pad) + str
}

func PadRight(str string, width int, pad string) string {
	return str + getPadString(width-len(str), pad)
}

func Pad(str string, width int, leftPad string, rightPad string) string {
	switch {
	case len(leftPad) == 0:
		return PadRight(str, width, rightPad)
	case len(rightPad) == 0:
		return PadLeft(str, width, leftPad)
	}
	padLen := (width - len(str)) / 2
	return getPadString(padLen, leftPad) + str + getPadString(width-len(str)-padLen, rightPad)
}

func Center(str string, width int) string {
	return Pad(str, width, " ", " ")
}

func MapLines(str string, fn func(string) string) string {
	arr := strings.Split(str, "\n")
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i])
	}
	return strings.Join(arr, "\n")
}

func AlignLeft(str string) string {
	return MapLines(str, func(line string) string {
		return strings.TrimLeft(line, " ")
	})
}

func AlignRight(str string, width int) string {
	return MapLines(str, func(line string) string {
		line = strings.Trim(line, " ")
		if line == "" {
			return ""
		}
		return PadLeft(line, width, " ")
	})
}

func AlignCenter(str string, width int) string {
	return MapLines(str, func(line string) string {
		line = strings.Trim(line, " ")
		if line == "" {
			return ""
		}
		return Center(line, width)
	})
}

const (
	AlignTypeCenter = "center"
	AlignTypeLeft   = "left"
	AlignTypeRight  = "right"
)

func Align(str string, typ string, width int) string {
	switch typ {
	case AlignTypeCenter:
		return AlignCenter(str, width)
	case AlignTypeRight:
		return AlignRight(str, width)
	default:
		return AlignLeft(str)
	}
}

func ExpandTabs(str string, count int) string {
	return strings.Replace(str, "\t", strings.Repeat(" ", count), -1)
}

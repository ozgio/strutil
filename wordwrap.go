package strutil

import "strings"

var newLine = rune('\n')
var space = rune(' ')

// strBuffer is a special string builder, which can track rune count, written
// mainly for WordWrap
type strBuffer struct {
	//internal array to store string
	buf []byte
	//count of runes.
	length int
}

func (b *strBuffer) WriteString(str string, length int) *strBuffer {
	b.buf = append(b.buf, str...)
	b.length += length
	return b
}

func (b *strBuffer) WriteRune(r rune) *strBuffer {
	b.buf = append(b.buf, string(r)...)
	b.length++
	return b
}

func (b *strBuffer) AppendStrBuffer(n strBuffer) *strBuffer {
	b.buf = append(b.buf, n.buf...)
	b.length += n.length
	return b
}

func (b *strBuffer) NewLine() *strBuffer {
	b.WriteString("\n", 1)
	return b
}

func (b *strBuffer) Reset() *strBuffer {
	b.buf = b.buf[:0]
	b.length = 0
	return b
}

func (b *strBuffer) LeftTrim() *strBuffer {
	for i, c := range b.buf {
		if c != ' ' {
			b.buf = b.buf[i:]
			b.length -= i
			return b
		}
	}

	b.buf = b.buf[:0]
	b.length = 0
	return b
}

func (b *strBuffer) String() string {
	return string(b.buf)
}

// Wordwrap wraps the given string str based on colLen as max line width.
// if breakLongWords is true, it breaks the words which are longer than colLen.
//
// Notes:
// - Wordwrap doesn't trim the lines, except it trims the left side of the new line
// created by breaking a long line.
// - Tabs should be converted to space before using WordWrap.
func Wordwrap(str string, colLen int, breakLongWords bool) string {
	if colLen == 0 || str == "" || len(str) < colLen {
		return str
	}
	var buff strings.Builder
	var line strBuffer
	var word strBuffer

	runeIndex := -1
	var runeWritten bool
	var r rune

	for _, r = range str {
		runeIndex++
		runeWritten = false

		if r == newLine {
			line.AppendStrBuffer(word).NewLine()
			buff.Write(line.buf)
			line.Reset()
			word.Reset()
			continue
		}
		if r == space {
			line.AppendStrBuffer(word)
			word.Reset()
		}

		if line.length+word.length+1 > colLen {
			if line.length > 0 {
				runeWritten = true
				word.WriteRune(r)
				line.NewLine()
				buff.Write(line.buf)
				line.Reset()
				word.LeftTrim()
			} else if breakLongWords {
				line.AppendStrBuffer(word).NewLine()
				buff.Write(line.buf)
				line.Reset()
				word.Reset()
			}
		}

		if !runeWritten {
			word.WriteRune(r)
		}
	}

	line.AppendStrBuffer(word)
	buff.Write(line.buf)

	res := buff.String()
	if r != newLine && len(res) > 0 && res[len(res)-1:] == "\n" {
		res = res[:len(res)-1]
	}

	return res
}

package strutil

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// UTF8Len is an alias of utf8.RuneCountInString which returns the number of
// runes in s. Erroneous and short encodings are treated as single runes of
// width 1 byte.
var UTF8Len = utf8.RuneCountInString

// Substring gets a part of the string between start and end. If end is 0,
// end is taken as the length of the string.
//
// It is UTF8 safe version of using slice notations in strings. It panics
// when the indexes are out of range. String length can be get with
// UTF8Len function before using Substring
func Substring(str string, start int, end int) string {
	runes := []rune(str)
	size := len(runes)
	if start < 0 || start > size-1 {
		panic(fmt.Sprintf("start (%d) is out of range (%d)", start, size))
	}
	if end == 0 {
		return string(runes[start:])
	}

	if end <= start || end > size {
		panic(fmt.Sprintf("end (%d) is out of range (%d)", end, size))

	}
	return string(runes[start:end])
}

// IsASCII checks if all the characters in string are in standard ASCII table
// It is taken from strings.Fields function
func IsASCII(s string) bool {
	// setBits is used to track which bits are set in the bytes of s.
	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
		setBits |= s[i]
	}

	return setBits < utf8.RuneSelf
}

// CountWords count the words, It uses the same base function with 'Words'
// function. only difference is CountWords doesn't allocate an array so
// it is faster and more memory efficient
func CountWords(str string) int {
	_, count := words(str, true)
	return count
}

// Words returns the words inside the text.
// - Numbers are counted as words
// - If they are inside a word these punctuations don't break a word: ', -, _
func Words(str string) []string {
	arr, _ := words(str, false)
	return arr
}

const (
	wordRune = iota
	wordPuncRune
	nonWordRune
)

// wordPuncRunes are punctuations which can be inside words: O'Neil, micro-service
var wordPuncRunes = [...]rune{rune('\''), rune('-'), rune('_')}

func inWordPuncRune(r rune) bool {
	for _, p := range wordPuncRunes {
		if r == p {
			return true
		}
	}
	return false
}

// words is the ugly base function for Word and CountWords. It returns words
// and the count of the words. If onlyCount is true only count is returned,
// no array is created.
func words(str string, onlyCount bool) ([]string, int) {
	var arr []string
	if !onlyCount {
		arr = make([]string, 0, len(str)/4) //TODO search for better start size
	}
	var prevCat = nonWordRune
	var lastStart = -1
	var count = 0

	for i, r := range str {
		var cat int
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			cat = wordRune
		case inWordPuncRune(r):
			//faster: case r == wordPuncRunes[0] || r == wordPuncRunes[1] || r == wordPuncRunes[2]:
			cat = wordPuncRune
		default:
			cat = nonWordRune
		}

		switch {
		//start word
		case cat == wordRune && prevCat != wordRune && lastStart == -1:
			lastStart = i
		//end word
		case cat == nonWordRune && (prevCat == wordRune || prevCat == wordPuncRune) && lastStart >= 0:
			if !onlyCount {
				arr = append(arr, str[lastStart:i])
			}
			lastStart = -1
			count++
		}

		prevCat = cat
	}

	if lastStart >= 0 {
		if !onlyCount {
			arr = append(arr, str[lastStart:])
		}
		count++
	}
	return arr, count
}

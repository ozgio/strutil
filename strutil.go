package strutil

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Reverse reverses the string
func Reverse(str string) string {
	runes := []rune(str)
	l := len(runes)
	for i := 0; i < l/2; i++ {
		runes[i], runes[l-i-1] = runes[l-i-1], runes[i]
	}
	return string(runes)
}

// ReplaceAllToOne replaces every string in the from to the string "to"
func ReplaceAllToOne(str string, from []string, to string) string {
	arr := make([]string, len(from)*2)
	for i, s := range from {
		arr[i*2] = s
		arr[i*2+1] = to
	}
	r := strings.NewReplacer(arr...)

	return r.Replace(str)
}

// SpecialAccentReplacer is a string.Replacer for removing accents for special
// characters like Turkish "ı" or "İ"
var SpecialAccentReplacer = strings.NewReplacer(
	"ı", "i",
	"İ", "I",
	"ð", "o",
	"ø", "o",
	"Ø", "O",
	"ß", "ss",
	"ł", "l",
	"æ", "a")

// RemoveAccents removes accents from the letters. The resuting string only has
// the letters from English alphabet.
// taken from https://blog.golang.org/normalization
func RemoveAccents(str string) (string, int, error) {
	str = SpecialAccentReplacer.Replace(str)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	return transform.String(t, str)
}

// CountWords count the words approximately based on strings.Fields function.
func CountWords(str string) int {
	var count int
	arr := strings.Fields(str)
	count = len(arr)

	return count
}

// Substring gets a part of the string between start and end. If end is -1,
// end is taken as the length of the string.
//
// It is UTF8 safe version of using slice notations in strings. It panics
// when the indexes are out of range. String length can be get with
// UTF8Len function before using Substring
func Substring(str string, start int, end int) string {
	runes := []rune(str)
	if start < 0 || start > len(runes)-1 {
		panic(fmt.Sprintf("start (%d) is out of range (%d)", start, len(runes)))
	}
	if end == -1 {
		return string(runes[start:])
	}

	if end <= start || end > len(runes) {
		panic(fmt.Sprintf("end (%d) is out of range (%d)", end, len(runes)))

	}
	return string(runes[start:end])
}

// UTF8Len is a alias of utf8.RuneCountInString which returns the number of
// runes in s. Erroneous and short encodings are treated as single runes of
// width 1 byte.
var UTF8Len = utf8.RuneCountInString

// MapLines runs function fn on every line of the string.
// It splits the string by new line "\n" and runs the fn for every line and
// returns the new string by combining these lines with "\n"
func MapLines(str string, fn func(string) string) string {
	arr := strings.Split(str, "\n")
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i])
	}
	return strings.Join(arr, "\n")
}

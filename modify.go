package strutil

import (
	"fmt"
	"strings"
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
// It may not be work as expected for some specific letters. Please create an
// issue for these situations.
func RemoveAccents(str string) (string, int, error) {
	str = SpecialAccentReplacer.Replace(str)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	return transform.String(t, str)
}

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

// Splice insert a new string in place of the string between start and end indexes.Splice
// It is based on runes so start and end indexes are rune based indexes.
// It can be used to remove a part of string by giving newStr as empty string
func Splice(str string, newStr string, start int, end int) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	size := len(runes)
	if start < 0 || start > size-1 {
		panic(fmt.Sprintf("start (%d) is out of range (%d)", start, size))
	}
	if end <= start || end > size {
		panic(fmt.Sprintf("end (%d) is out of range (%d)", end, size))
	}
	return string(runes[:start]) + newStr + string(runes[end:])
}

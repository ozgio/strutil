package strutil

import (
	"fmt"
	"strings"

	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Reverse(str string) string {
	runes := []rune(str)
	l := len(runes)
	for i := 0; i < l/2; i++ {
		runes[i], runes[l-i-1] = runes[l-i-1], runes[i]
	}
	return string(runes)
}

func ReplaceAllToOne(str string, from []string, to string) string {
	arr := make([]string, len(from)*2)
	for i, s := range from {
		arr[i*2] = s
		arr[i*2+1] = to
	}
	r := strings.NewReplacer(arr...)

	return r.Replace(str)
}

//SpecialAccentReplacer is string.Replacer for removing accents for special
//characters like Turkish "ı" or "İ"
var SpecialAccentReplacer = strings.NewReplacer("ı", "i", "İ", "I")

//RemoveAccents removes accents from the letters. The resuting string only has the letters from English alphabet
//taken from https://blog.golang.org/normalization
func RemoveAccents(str string) (string, int, error) {
	str = SpecialAccentReplacer.Replace(str)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	return transform.String(t, str)
}

// const (
// 	runeLetter = iota
// 	runeNumber
// 	runePunct
// )

// var wordsRegexp *regexp.Regexp = regexp.MustCompile(`[\p{L}\d-']+`)
// var excludeWordsRegexp *regexp.Regexp = regexp.MustCompile(`^[-']+\s|\s[-']+\s|\s[-']+$`)

// CountWords count the words approximately based on strings.Fields function.
func CountWords(str string) int {
	var count int
	arr := strings.Fields(str)
	count = len(arr)

	return count
}

func UnicodeSubstring(str string, start int, end int) string {
	runes := []rune(str)
	if start < 0 || start > len(runes)-1 {
		panic(fmt.Sprintf("start (%d) is out of range (%d)", start, len(runes)))
	}
	if end <= start || end > len(runes) {
		panic(fmt.Sprintf("end (%d) is out of range (%d)", end, len(runes)))

	}
	return string(runes[start:end])
}

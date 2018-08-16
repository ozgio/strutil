package strutil

import (
	"regexp"
	"strings"
	"unicode"
)

// ToSnakeCase converts str into snake_case formatted string. In the process it
// trims the string and the converts characters into lowercase. Only space " "
// character is converted into underscore "_". If you have other characters
// you should convert them into spaces before calling ToSnakeCase
//
// Example:
//     ToSnakeCase("Snake Case")
//     //Outputs: snake_case
func ToSnakeCase(str string) string {
	str = strings.TrimSpace(strings.ToLower(str))
	return strings.Replace(str, " ", "_", -1)
}

// ToCamelCase converts str into camelCase formatted string after trimming it.
// It doesn't change the cases of letters except the first letters of the words.
// ToCamelCase also doesn't remove punctions or such characters and it seperates
// words only with " "
//
// Example:
//     ToCamelCase("camel case")
//     //Outputs: camelCase
// 	   ToCamelCase("inside dynaMIC-HTML")
// 	   //Outputs: insideDynaMIC-HTML
func ToCamelCase(str string) string {
	str = strings.TrimSpace(str)
	if UTF8Len(str) < 2 {
		return str
	}

	var buff strings.Builder
	var prev string
	for _, r := range str {
		c := string(r)
		if c != " " {
			if prev == " " {
				c = strings.ToUpper(c)
			}
			buff.WriteString(c)
		}
		prev = c
	}
	return buff.String()
}

// SplitCamelCase splits and returns words in camelCase format.
//
// Example:
//     SplitCamelCase("loremIpsum")
//	   //Output
//     {"lorem", "ipsum"}
func SplitCamelCase(str string) []string {
	str = strings.TrimSpace(str)
	if UTF8Len(str) < 2 {
		return []string{str}
	}
	var prev rune
	var start int
	words := []string{}
	runes := []rune(str)

	for i, r := range runes {
		if i != 0 {
			switch {
			case unicode.IsLetter(r) && unicode.IsDigit(r):
				fallthrough
			case unicode.IsDigit(r) && unicode.IsLetter(prev):
				fallthrough
			case unicode.IsUpper(r) && unicode.IsLower(prev):
				words = append(words, string(runes[start:i]))
				start = i
			case unicode.IsLower(r) && unicode.IsUpper(prev) && start != i-1:
				words = append(words, string(runes[start:i-1]))
				start = i - 1
			}
		}
		prev = r
	}

	if start < len(runes) {
		words = append(words, string(runes[start:]))
	}
	return words
}

// nonAlphnumRegex is a regexp for selecting non alphanumeric characters
var nonAlphnumRegex *regexp.Regexp = regexp.MustCompile("[^a-z0-9]")

// multiSpaceRegex is a regexp for select contigous spaces
var multiSpaceRegex *regexp.Regexp = regexp.MustCompile("[ ]{2,}")

// Slugify converts a string to a slug which is useful in URLs, filenames.
// It removes accents, converts to lower case, remove the characters which
// are not letters or numbers and replaces spaces with "-".
func Slugify(str string) string {
	return SlugifySpecial(str, "-")
}

// SlugifySpecial converts a string to a slug with the delimeter.
// It removes accents, converts string to lower case, remove the characters
// which are not letters or numbers and replaces spaces with the delimeter.
func SlugifySpecial(str string, delimeter string) string {
	str, _, err := RemoveAccents(str)
	if err != nil {
		return ""
	}
	str = strings.ToLower(str)
	str = nonAlphnumRegex.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	str = multiSpaceRegex.ReplaceAllString(str, " ")
	str = strings.Replace(str, " ", delimeter, -1)

	return str
}

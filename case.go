package strutil

import (
	"strings"
	"unicode"
)

// ToSnakeCase converts string into snake_case formatted string. In the process it
// trims the string and then converts characters into lowercase. Only space " "
// character is converted into underscore "_". If you have other characters
// you should convert them into spaces before calling ToSnakeCase
//
// Example:
//     ToSnakeCase("Snake Case") //Output: snake_case
func ToSnakeCase(str string) string {
	str = strings.TrimSpace(strings.ToLower(str))
	return strings.Replace(str, " ", "_", -1)
}

// ToCamelCase converts string into camelCase formatted string after trimming it.
// It doesn't change the cases of letters except the first letters of the words.
// ToCamelCase also doesn't remove punctions or such characters and it separates
// words only with " "
//
// Example:
//     ToCamelCase("camel case") //Output: camelCase
// 	   ToCamelCase("inside dynaMIC-HTML") //Output: insideDynaMIC-HTML
func ToCamelCase(str string) string {
	str = strings.TrimSpace(str)
	if Len(str) < 2 {
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
//   SplitCamelCase("loremIpsum") //Output []string{"lorem", "Ipsum"}
func SplitCamelCase(str string) []string {
	str = strings.TrimSpace(str)
	if Len(str) < 2 {
		return []string{str}
	}
	var prev rune
	var start int
	words := []string{}
	runes := []rune(str)

	for i, r := range runes {
		if i != 0 {
			switch {
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

package strutil

import (
	"strings"
	"unicode"
)

func ToSnakeCase(str string) string {
	str = strings.TrimSpace(strings.ToLower(str))
	return strings.Replace(str, " ", "_", -1)
}

func ToCamelCase(str string) string {
	str = strings.TrimSpace(str)
	if len(str) < 2 {
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

func SplitCamelCase(str string) []string {
	str = strings.TrimSpace(str)
	if len(str) < 2 {
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

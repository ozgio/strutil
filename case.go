package strutil

import (
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

	delBytes := []byte(delimeter)

	n := make([]byte, 0, len(str))
	isPrevSpace := false
	for _, r := range str {
		//toLowerCase
		if r >= 'A' && r <= 'Z' {
			r -= 'A' - 'a'
		}

		//replace non-alphanum chars with delimeter
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			n = append(n, byte(int8(r)))
			isPrevSpace = false
		case !isPrevSpace:
			n = append(n, delBytes...)
			fallthrough
		default:
			isPrevSpace = true
		}
	}

	//trim right
	ln := len(n)
	ld := len(delimeter)
	if ln >= ld && string(n[ln-ld:]) == delimeter {
		n = n[:ln-ld]
	}

	return string(n)
}

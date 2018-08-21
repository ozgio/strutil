package strutil

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

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

// Slugify converts a string to a slug which is useful in URLs, filenames.
// It removes accents, converts to lower case, remove the characters which
// are not letters or numbers and replaces spaces with "-".
//
// Currently it doesn't support non-latin alphabets. See RemoveAccents for
// latin alphabet support
func Slugify(str string) string {
	return SlugifySpecial(str, "-")
}

// SlugifySpecial converts a string to a slug with the delimeter.
// It removes accents, converts string to lower case, remove the characters
// which are not letters or numbers and replaces spaces with the delimeter.
//
// Currently it doesn't support non-latin alphabets. See RemoveAccents for
// latin alphabet support
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

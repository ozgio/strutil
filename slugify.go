package strutil

// Slugify converts a string to a slug which is useful in URLs, filenames.
// It removes accents, converts to lower case, remove the characters which
// are not letters or numbers and replaces spaces with "-".
//
// Example:
//   strutil.Slugify("'We löve Motörhead'") //Output: we-love-motorhead
//
// Normalzation is done with strutil.ReplaceAccents function using a rune replacement map
// You can use the following code for better normalization before strutil.Slugify()
//
//   str := "'We löve Motörhead'"
//   t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
//   str = transform.String(t, str) //We love Motorhead
//
// Slugify doesn't support transliteration. You should use a transliteration
// library before Slugify like github.com/rainycape/unidecode
//
// Example:
//   import "github.com/rainycape/unidecode"
//
//   str := unidecode.Unidecode("你好, world!")
//   strutil.Slugify(str) //Output: ni-hao-world
//
func Slugify(str string) string {
	return SlugifySpecial(str, "-")
}

// SlugifySpecial converts a string to a slug with the delimiter.
// It removes accents, converts string to lower case, remove the characters
// which are not letters or numbers and replaces spaces with the delimiter.
//
// Example:
//   strutil.SlugifySpecial("'We löve Motörhead'", "-") //Output: we-love-motorhead
//
// SlugifySpecial doesn't support transliteration. You should use a transliteration
// library before SlugifySpecial like github.com/rainycape/unidecode
//
// Example:
//   import "github.com/rainycape/unidecode"
//
//   str := unidecode.Unidecode("你好, world!")
//   strutil.SlugifySpecial(str, "-") //Output: ni-hao-world
func SlugifySpecial(str string, delimiter string) string {
	str = RemoveAccents(str)

	delBytes := []byte(delimiter)

	n := make([]byte, 0, len(str))
	isPrevSpace := false
	for _, r := range str {
		//toLowerCase
		if r >= 'A' && r <= 'Z' {
			r -= 'A' - 'a'
		}

		//replace non-alphanum chars with delimiter
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			n = append(n, byte(int8(r)))
			isPrevSpace = false
		case !isPrevSpace:
			if len(n) > 0 {
				n = append(n, delBytes...)
			}
			fallthrough
		default:
			isPrevSpace = true
		}
	}

	//trim right
	ln := len(n)
	ld := len(delimiter)
	if ln >= ld && string(n[ln-ld:]) == delimiter {
		n = n[:ln-ld]
	}

	return string(n)
}

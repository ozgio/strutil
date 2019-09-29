package strutil

import "unicode/utf8"

// Len is an alias of utf8.RuneCountInString which returns the number of
// runes in s. Erroneous and short encodings are treated as single runes of
// width 1 byte.
var Len = utf8.RuneCountInString

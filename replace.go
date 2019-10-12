package strutil

import (
	"fmt"
	"strings"
)

// ReplaceAllToOne replaces every string in the "from" with the string "to"
func ReplaceAllToOne(str string, from []string, to string) string {
	arr := make([]string, len(from)*2)
	for i, s := range from {
		arr[i*2] = s
		arr[i*2+1] = to
	}
	r := strings.NewReplacer(arr...)

	return r.Replace(str)
}

// Splice insert a new string in place of the string between start and end indexes.
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

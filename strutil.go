package strutil

import (
	"strings"
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

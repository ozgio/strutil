package strutil

import "strings"

func StrBefore(str string, search string) string {
	pos := strings.Index(str, search)
	if pos == -1 {
		return ""
	}
	return strings.Trim(str[0:pos], " ")
}

func StrBeforeLast(str string, search string) string {
	pos := strings.LastIndex(str, search)
	if pos == -1 {
		return ""
	}
	return strings.Trim(str[0:pos], " ")
}

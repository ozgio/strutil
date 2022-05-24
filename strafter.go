package strutil

import "strings"

func StrAfter(str string, search string) string {
	pos := strings.Index(str, search)
	return StrAfterIndex(str, search, pos)
}

func StrAfterLast(str string, search string) string {
	pos := strings.LastIndex(str, search)
	return StrAfterIndex(str, search, pos)
}

func StrAfterIndex(str string, search string, pos int) string {
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(search)
	if adjustedPos >= len(str) {
		return ""
	}
	return strings.Trim(str[adjustedPos:len(str)], " ")
}

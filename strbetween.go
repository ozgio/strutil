package strutil

import "strings"

func StrBetween(str string, a string, b string) string {
	posFirst := strings.Index(str, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(str, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return strings.Trim(str[posFirstAdjusted:posLast], " ")
}

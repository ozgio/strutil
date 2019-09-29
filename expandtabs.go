package strutil

import "strings"

//ExpandTabs convert tabs the spaces with the length of count
func ExpandTabs(str string, count int) string {
	return strings.Replace(str, "\t", strings.Repeat(" ", count), -1)
}

package strutil

import "strings"

// Indent indents every line of string str with the left parameter
// Empty lines are indented too.
func Indent(str string, left string) string {
	return left + strings.Replace(str, "\n", "\n"+left, -1)
}

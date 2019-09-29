package strutil

import "strings"

// MapLines runs function fn on every line of the string.
// It splits the string by new line "\n" and runs the fn for every line and
// returns the new string by combining these lines with "\n"
func MapLines(str string, fn func(string) string) string {
	arr := strings.Split(str, "\n")
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i])
	}
	return strings.Join(arr, "\n")
}

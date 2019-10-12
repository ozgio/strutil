package strutil

import "strings"

// MapLines runs function fn on every line of the string.
// It splits the string by new line character ("\n"), then runs 'fn' for every line and
// returns the new string by combining these lines with "\n"
func MapLines(str string, fn func(string) string) string {
	return SplitAndMap(str, "\n", fn)
}

//SplitAndMap splits the string and runs the function fn on every part
func SplitAndMap(str string, split string, fn func(string) string) string {
	arr := strings.Split(str, split)
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i])
	}
	return strings.Join(arr, split)
}

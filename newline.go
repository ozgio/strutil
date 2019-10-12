package strutil

import "runtime"

// OSNewLine returns operating system's default new line character. It is \r\n
// in Windowns and \n elsewhere.
func OSNewLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

package strutil

import "runtime"

func DefaultNewLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

var defaultNewLine = DefaultNewLine()

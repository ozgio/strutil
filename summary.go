package strutil

import "strings"

//Summary cuts the string to a new length and adds the "end" to it
//It only break up the words by spaces
func Summary(str string, length int, end string) string {
	if str == "" || length <= 0 {
		return str
	}
	var runeIndex int
	var i int
	var r rune
	var lastSpaceIndex int
	for i, r = range str {
		switch {
		case r == newLine:
			return str[:i] + end
		case r == ' ':
			lastSpaceIndex = i
		}

		if runeIndex+1 > length {
			break
		}
		runeIndex++
	}

	if runeIndex < length {
		return str
	}

	if lastSpaceIndex == 0 {
		//no space until here, so break the word
		str = str[:i]
	} else {
		//break from the last seen space
		str = str[:lastSpaceIndex]
		str = strings.TrimRight(str, " ")
	}

	return str + end
}

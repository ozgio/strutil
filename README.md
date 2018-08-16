String Utilities for Go
=======================
[![CircleCI](https://circleci.com/gh/ozgio/strutil.svg?style=svg)](https://circleci.com/gh/ozgio/strutil)

"strutil" provides UTF8 safe string functions for go applications. 

For documentation with examples see [GoDoc](https://godoc.org/github.com/ozgio/strutil)

## Functions

| Function          | Desctiption                                           |
|-------------------|-------------------------------------------------------|
| Align             | Aligns the text to spesified side                     |
| AlignCenter       | Aligns the text to the center                         |
| AlignLeft         | Aligns the text to the left                           |
| AlignRight        | Aligns the text to the right                          |
| Box               | Draws a frame around the string with default chars    |
| Center            | Centers the string                                    |
| CountWords        | Count the words in the string                         |
| CustomBox         | Draws a frame aroud the string with spesified chars   |
| ExpandTabs        | Converts tabs to spaces                               |
| Indent            | Indents the string                                    |
| MapLines          | Runs spesified function on every line of the text     |
| Pad               | Left and right pads the strings                       |
| PadLeft           | Left pads the string                                  |
| PadRight          | Right pads the string                                 |
| RemoveAccents     | Remove accents in the string                          |
| ReplaceAllToOne   | Replace all strings in a text to a spesified string   |
| Reverse           | Reverses the string                                   |
| Slugify           | Converts the string to a slug                         |
| SplitCamelCase    | Splits the words in a camelCase string                |
| ToCamelCase       | Converts the string to camelCase                      |
| ToSnakeCase       | Converts the string to snake_Case                     |
| Substring         | Returns a part of the string                          |
| Wordwrap          | Wraps the lines in the text                           |

## Install 

Prequsities:
- Go 1.10+

Install with 

    go get github.com/ozgio/strutil

Dependencies:

    go get golang.org/x/text/runes
	go get golang.org/x/text/transform
	go get golang.org/x/text/unicode/norm

Import

    import "github.com/ozgio/strutil"

## TODO
- Fix problems in some edge cases
- Optimize and improve code
- CI Integration
- Add functions:
  - Splice
  - Truncate
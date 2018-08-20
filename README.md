String Utilities for Go
=======================
[![CircleCI](https://circleci.com/gh/ozgio/strutil.svg?style=svg)](https://circleci.com/gh/ozgio/strutil) 
[![GoReportCard](https://goreportcard.com/badge/github.com/ozgio/strutil)](https://goreportcard.com/report/github.com/ozgio/strutil)
[![GoDocs](https://godoc.org/github.com/ozgio/strutil?status.svg)](https://godoc.org/github.com/ozgio/strutil)


"strutil" provides fast, UTF8 safe string functions for go applications. 

For documentation with examples see [GoDoc](https://godoc.org/github.com/ozgio/strutil)

## Functions

| Function                                                                  | Desctiption                                           |
|---------------------------------------------------------------------------|-------------------------------------------------------|
| [Align](https://godoc.org/github.com/ozgio/strutil#Align)                 | Aligns the text to the spesified side                 |
| [AlignCenter](https://godoc.org/github.com/ozgio/strutil#AlignCenter)     | Aligns the text to the center                         |
| [AlignLeft](https://godoc.org/github.com/ozgio/strutil#AlignLeft)         | Aligns the text to the left                           |
| [AlignRight](https://godoc.org/github.com/ozgio/strutil#AlignRight)       | Aligns the text to the right                          |
| [Box](https://godoc.org/github.com/ozgio/strutil#Box)                     | Draws a frame around the string with default chars    |
| [Center](https://godoc.org/github.com/ozgio/strutil#Center)               | Centers the string                                    |
| [CountWords](https://godoc.org/github.com/ozgio/strutil#CountWords)       | Count the words in the string                         |
| [CustomBox](https://godoc.org/github.com/ozgio/strutil#CustomBox)         | Draws a frame aroud the string with spesified chars   |
| [ExpandTabs](https://godoc.org/github.com/ozgio/strutil#ExpandTabs)       | Converts tabs to spaces                               |
| [Indent](https://godoc.org/github.com/ozgio/strutil#Indent)               | Indents the string                                    |
| [MapLines](https://godoc.org/github.com/ozgio/strutil#MapLines)           | Runs spesified function on every line of the text     |
| [Pad](https://godoc.org/github.com/ozgio/strutil#Pad)                     | Left and right pads the strings                       |
| [PadLeft](https://godoc.org/github.com/ozgio/strutil#PadLeft)             | Left pads the string                                  |
| [PadRight](https://godoc.org/github.com/ozgio/strutil#PadRight)           | Right pads the string                                 |
| [Random](https://godoc.org/github.com/ozgio/strutil#Random)               | Creates a random string from a character set          |
| [RemoveAccents](https://godoc.org/github.com/ozgio/strutil#RemoveAccents) | Remove accents in the string                          |
| [ReplaceAllToOne](https://godoc.org/github.com/ozgio/strutil#ReplaceAllToOne) | Replace all substrings in the text to the spesified string   |
| [Reverse](https://godoc.org/github.com/ozgio/strutil#Reverse)             | Reverses the string                                   |
| [Slugify](https://godoc.org/github.com/ozgio/strutil#Slugify)             | Converts the string to a slug                         |
| [SplitCamelCase](https://godoc.org/github.com/ozgio/strutil#SplitCamelCase)   | Splits the words in a camelCase string            |
| [ToCamelCase](https://godoc.org/github.com/ozgio/strutil#ToCamelCase)     | Converts the string to camelCase                      |
| [ToSnakeCase](https://godoc.org/github.com/ozgio/strutil#ToSnakeCase)     | Converts the string to snake_Case                     |
| [Substring](https://godoc.org/github.com/ozgio/strutil#Substring)         | Returns a part of the string                          |
| [Summary](https://godoc.org/github.com/ozgio/strutil#Summary)             | Cuts the text to the length                           |
| [Splice](https://godoc.org/github.com/ozgio/strutil#Splice)               | Replaces a part of the string                         |
| [Wordwrap](https://godoc.org/github.com/ozgio/strutil#Wordwrap)           | Wraps the lines in the text                           |

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
- Improve tests. More test cases are needed
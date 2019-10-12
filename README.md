String Utilities for Go
=======================
[![CircleCI](https://circleci.com/gh/ozgio/strutil.svg?style=svg)](https://circleci.com/gh/ozgio/strutil)
[![GoReportCard](https://goreportcard.com/badge/github.com/ozgio/strutil)](https://goreportcard.com/report/github.com/ozgio/strutil)
[![GoDocs](https://godoc.org/github.com/ozgio/strutil?status.svg)](https://godoc.org/github.com/ozgio/strutil)

"strutil" provides string functions for go applications.

For documentation with examples see [GoDoc](https://godoc.org/github.com/ozgio/strutil)

## Functions


### Align ([Docs](https://godoc.org/github.com/ozgio/strutil#Align))
Aligns the text to the spesified side

```go
strutil.Align("lorem ipsum", strutil.Right, 20) //->"        lorem ipsum" 
```

### AlignCenter ([Docs](https://godoc.org/github.com/ozgio/strutil#AlignCenter))
Aligns the text to the center
```go
strutil.AlignCenter("lorem\nipsum", 10) //->"  lorem   \n  ipsum   " 
```

### AlignLeft ([Docs](https://godoc.org/github.com/ozgio/strutil#AlignLeft))
Aligns the text to the left
```go
strutil.AlignLeft("  lorem   \n  ipsum   ") //->"lorem   \nipsum   " 
```

### AlignRight ([Docs](https://godoc.org/github.com/ozgio/strutil#AlignRight))
Aligns the text to the right
```go
strutil.AlignRight("lorem\nipsum", 10) //-> "     lorem\n     ipsum" 
```

### CountWords ([Docs](https://godoc.org/github.com/ozgio/strutil#CountWords))
Counts the words
```go
strutil.CountWords("Lorem ipsum, dolor sit amet") //-> "5"
```

### DrawBox ([Docs](https://godoc.org/github.com/ozgio/strutil#DrawBox))
Draws a frame around the string with default character set
```go
strutil.DrawBox("Hello World", 20, strutil.Center)
//┌──────────────────┐
//│   Hello World    │
//└──────────────────┘
```

### DrawCustomBox ([Docs](https://godoc.org/github.com/ozgio/strutil#DrawCustomBox))
Draws a frame around the string with custom character set
```go
strutil.DrawCustomBox("Hello World", 20, strutil.Center, strutil.SimpleBox9Slice(), "\n")
//┌──────────────────┐
//│   Hello World    │
//└──────────────────┘
```

### ExpandTabs  ([Docs](https://godoc.org/github.com/ozgio/strutil#ExpandTabs))
Converts tabs to the spaces
```go
strutil.ExpandTabs("\tlorem\n\tipsum\n", 2) //-> "  lorem\n  ipsum\n"
```

### Indent ([Docs](https://godoc.org/github.com/ozgio/strutil#Indent))
Indents every line
```go
strutil.Indent("lorem\nipsum", "> ") //-> "> lorem\n> ipsum"
```

### IsASCII ([Docs](https://godoc.org/github.com/ozgio/strutil#IsASCII))
Checks if all the characters in string are in standard ASCII table
```go
strutil.IsASCII("lorem\nipsum") //-> true
```

### Len ([Docs](https://godoc.org/github.com/ozgio/strutil#Len))
Alias of utf8.RuneCountInString which returns the number of runes in string
```go
strutil.Len("böt") //-> "3"
```

### MapLines ([Docs](https://godoc.org/github.com/ozgio/strutil#MapLines))
Runs function fn on every line of the string
```go
strutil.MapLines("   lorem      \n    ipsum      ", strings.TrimSpace) //-> "lorem\nipsum"
```

### OSNewLine ([Docs](https://godoc.org/github.com/ozgio/strutil#OSNewLine))
OSNewLine returns operating systems default new line character
```go
strutil.OSNewLine() //-> "\n"
```

### Pad ([Docs](https://godoc.org/github.com/ozgio/strutil#Pad))
Left and right pads the string
```go
strutil.Pad("lorem", 11, "->", "<-") //-> "->->lorem<-<-"
```

### PadLeft ([Docs](https://godoc.org/github.com/ozgio/strutil#PadLeft))
Left pads the string
```go
strutil.PadLeft("lorem", 9, "->") //-> "->->lorem"
```

### PadRight ([Docs](https://godoc.org/github.com/ozgio/strutil#PadRight))
Right pads the string
```go
strutil.PadRight("lorem", 9, "<-") //-> "lorem<-<-"
```

### Random ([Docs](https://godoc.org/github.com/ozgio/strutil#Random))
Creates a random string from a character set
```go
strutil.Random("abcdefghi", 10) //-> "aciafbeafg"
```

### RemoveAccents ([Docs](https://godoc.org/github.com/ozgio/strutil#RemoveAccents))
Convert accented letters to ascii counterparts
```go
strutil.RemoveAccents("résumé") //-> "resume"
```

### ReplaceAllToOne ([Docs](https://godoc.org/github.com/ozgio/strutil#ReplaceAllToOne))
Replace all substrings in the text with the spesified string
```go
strutil.ReplaceAllToOne("lorem ipsum", []string{"o","e","i","u"}, ".") //-> "l.r.m .ps.m"
```

### Reverse ([Docs](https://godoc.org/github.com/ozgio/strutil#Reverse))
Reverses the string
```go
strutil.Reverse("lorem") //-> "merol"
```

### Splice ([Docs](https://godoc.org/github.com/ozgio/strutil#Splice))
Replaces a part of the string 
```go
strutil.Splice("lorem", "-x-", 2, 3) //-> "lo-x-em"
```

### SplitAndMap ([Docs](https://godoc.org/github.com/ozgio/strutil#SplitAndMap))
Splits the string and runs the function fn on every part
```go
strutil.MapLines("lorem-ipsum-dolor", "-",  strutil.Reverse) //-> "merol\nmuspi\nrolod"
```

### Slugify ([Docs](https://godoc.org/github.com/ozgio/strutil#Slugify))
Converts the string to a slug 
```go
strutil.Slugify("Lorem ipsum, dolör") //-> "lorem-ipsum-dolor"
```

### SlugifySpecial ([Docs](https://godoc.org/github.com/ozgio/strutil#SlugifySpecial))
Converts the string to a slug with custom delimeter.
```go
strutil.SlugifySpecial("Lorem ipsum, dolör", "_") //-> "lorem_ipsum_dolor"
```

### SplitCamelCase  ([Docs](https://godoc.org/github.com/ozgio/strutil#SplitCamelCase))
Splits the words in a camelCase string
```go
strutil.SplitCamelCase("loremIpsum") //-> []string{"lorem", "Ipsum"}
```

### Substring ([Docs](https://godoc.org/github.com/ozgio/strutil#SafeSubstring))
Gets a part of the string without panics
```go
strutil.SafeSubstring("lorem", 0, 1) //-> "l"
```

### MustSubstring ([Docs](https://godoc.org/github.com/ozgio/strutil#Substring))
Gets a part of the string
```go
strutil.Substring("lorem", 0, 1) //-> "l"
```

### Summary ([Docs](https://godoc.org/github.com/ozgio/strutil#Summary))
Cuts the string to the specified length
```go
strutil.Summary("Lorem ipsum dolor sit amet",  10, "...") //-> "lorem ipsum..."
```

### Tile ([Docs](https://godoc.org/github.com/ozgio/strutil#Tile))
Repeats the pattern until the result reaches the 'length'
```go
strutil.Tile("-৹", 4) //-> "-৹-৹"
```

### ToSnakeCase  ([Docs](https://godoc.org/github.com/ozgio/strutil#ToSnakeCase))
Converts the string to snake_case 
```go
strutil.ToSnakeCase("Snake Case") //-> "snake_case"
```

### ToCamelCase  ([Docs](https://godoc.org/github.com/ozgio/strutil#ToCamelCase))
Converts the string to camelCase 
```go
strutil.ToCamelCase("Camel Case") //-> "camelCase"
```

### Words ([Docs](https://godoc.org/github.com/ozgio/strutil#Words))
Returns the words inside the text
```go
strutil.Words("Lorem ipsum, dolor sit amet") //-> []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
```

### WordWrap ([Docs](https://godoc.org/github.com/ozgio/strutil#WordWrap))
Wraps the lines in the text  
```go
strutil.WordWrap("Lorem ipsum dolor sit amet", 15, false) //-> "Lorem ipsum\ndolor sit amet"
```


## Install

Prequsities:
- Go 1.10+

Install with
```sh
go get github.com/ozgio/strutil
```

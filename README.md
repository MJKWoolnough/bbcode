# bbcode

[![CI](https://github.com/MJKWoolnough/bbcode/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/bbcode/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/bbcode.svg)](https://pkg.go.dev/vimagination.zapto.org/bbcode)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/bbcode)](https://goreportcard.com/report/vimagination.zapto.org/bbcode)

--
    import "vimagination.zapto.org/bbcode"

Package bbcode implements a bbcode parser and converter.

## Highlights

 - Parse BBCode to HTML or plain text.
 - Allows for custom tags.
 - Filter attributes for safe rendering.

## Usage

```go
package main

import (
	"os"

	"vimagination.zapto.org/bbcode"
	"vimagination.zapto.org/bbcode/bbcodehtml"
)

func main() {
	parser := bbcode.New(bbcodehtml.All...)

	parser.ConvertString(os.Stdout, `This is [b]Bold[/b], this is [i]Italic[/i], and this is a [url=http://www.example.com]link[/url].`)

	// Output:
	// This is <b>Bold</b>, this is <i>Italic</i>, and this is a <a href="http://www.example.com">link</a>.
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/bbcode

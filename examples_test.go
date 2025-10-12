package bbcode_test

import (
	"os"

	"vimagination.zapto.org/bbcode"
	"vimagination.zapto.org/bbcode/bbcodehtml"
)

func Example() {
	parser := bbcode.New(bbcodehtml.All...)

	parser.ConvertString(os.Stdout, `This is [b]Bold[/b], this is [i]Italic[/i], and this is a [url=http://www.example.com]link[/url].`)

	// Output:
	// This is <b>Bold</b>, this is <i>Italic</i>, and this is a <a href="http://www.example.com">link</a>.
}

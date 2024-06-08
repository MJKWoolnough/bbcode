// Package bbcode implements a bbcode parser and converter
package bbcode // import "vimagination.zapto.org/bbcode"

import (
	"io"

	"vimagination.zapto.org/parser"
)

// Config changes how the syntax of the inteperated bbCode.
type Config struct {
	// TagOpen is the character used to open the tags.
	// In the default configuration this is the open bracket '['.
	TagOpen rune

	// TagClose is the character used to close the tags.
	// In the default configuration this is the close bracket ']'.
	TagClose rune

	// ClosingTag is the character used to define a closing tag, as opposed
	// to an opening tag.
	// In the default configuration, this is the slash '/'.
	ClosingTag rune

	// AttributeSep is the character used to separate the tag name from the
	// attribute.
	// In the default configuration this is the equals sign '='.
	AttributeSep rune

	// ValidTagName lists the characters that are allowed in the tag names.
	// Neither of the TagClose or AttributeSep characters should be in this
	// list.
	// In the default configuration this is A-Z a-z 0-9 and the asterix.
	ValidTagName string
}

var defaultConfig = Config{
	TagOpen:      '[',
	TagClose:     ']',
	ClosingTag:   '/',
	AttributeSep: '=',
	ValidTagName: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789*",
}

// BBCode is a fully configured parser.
type BBCode struct {
	tks  tokeniser
	text Handler
	tags []Handler
}

// NewWithConfig creates a bbCode parser with a custom bbCode configuration.
// The tags are a list of Handlers that will be used to process the tag tokens
// that are generated by the parser.
// The first Handler with an empty string for a name with be used to process
// the text. By default, this is the HTMLText handler.
func NewWithConfig(c Config, tags ...Handler) *BBCode {
	var text Handler = HTMLText

	for _, tag := range tags {
		if tag.Name() == "" {
			text = tag

			break
		}
	}

	return &BBCode{
		tks:  getTokeniser(c),
		text: text,
		tags: tags,
	}
}

// New create a new bbCode parser with the default configuration.
// See NewWithConfig for more information.
func New(tags ...Handler) *BBCode {
	return NewWithConfig(defaultConfig, tags...)
}

// Convert parses the byte slice and writes the output to the writer.
// Any error will be from the writing process and not from the parser.
func (b *BBCode) Convert(w io.Writer, input []byte) error {
	return b.convert(w, parser.NewByteTokeniser(input))
}

// ConvertString functions like Convert, but takes a string.
func (b *BBCode) ConvertString(w io.Writer, input string) error {
	return b.convert(w, parser.NewStringTokeniser(input))
}

// ConvertReader functions like Convert, but takes a io.Reader.
func (b *BBCode) ConvertReader(w io.Writer, input io.Reader) error {
	return b.convert(w, parser.NewReaderTokeniser(input))
}

func (b *BBCode) convert(w io.Writer, t parser.Tokeniser) error {
	p := Processor{
		w:      w,
		p:      b.tks.getParser(t),
		bbCode: b,
	}

	p.Process("")

	return p.err
}

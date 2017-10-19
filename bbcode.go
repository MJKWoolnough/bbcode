// Package bbcode implements a bbcode parser and converter
package bbcode

import (
	"io"

	"github.com/MJKWoolnough/parser"
)

/*
type Config struct {
	TagOpen, TagClose, ClosingTag, AttributeSep rune
	Name                                        string
}

var DefaultConfig = Config{
	TagOpen:      '[',
	TagClose:     ']',
	ClosingTag:   '/',
	AttributeSep: '=',
	Name:         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789*",
}
*/
type BBCode struct {
	//config Config
	text Handler
	tags []Handler
}

func New( /*c Config,*/ tags ...Handler) *BBCode {
	var text Handler = PlainText
	for _, tag := range tags {
		if tag.Name() == "" {
			text = tag
			break
		}
	}
	return &BBCode{
		//config: c,
		text: text,
		tags: tags,
	}
}

func (b *BBCode) Convert(w io.Writer, input []byte) error {
	return b.convert(w, parser.NewByteTokeniser(input))
}

func (b *BBCode) ConvertString(w io.Writer, input string) error {
	return b.convert(w, parser.NewStringTokeniser(input))
}

func (b *BBCode) ConvertReader(w io.Writer, input io.Reader) error {
	return b.convert(w, parser.NewReaderTokeniser(input))
}

func (b *BBCode) convert(w io.Writer, t parser.Tokeniser) error {
	p := Processor{
		w:    w,
		p:    newTokeniser(t),
		text: b.text,
		tags: b.tags,
	}
	p.Process("")
	return p.err
}

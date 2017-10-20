// Package bbcode implements a bbcode parser and converter
package bbcode

import (
	"io"

	"github.com/MJKWoolnough/parser"
)

type Config struct {
	TagOpen, TagClose, ClosingTag, AttributeSep rune
	Name                                        string
}

var defaultConfig = Config{
	TagOpen:      '[',
	TagClose:     ']',
	ClosingTag:   '/',
	AttributeSep: '=',
	Name:         "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789*",
}

type BBCode struct {
	tks  tokeniser
	text Handler
	tags []Handler
}

func NewWithConfig(c Config, tags ...Handler) *BBCode {
	var text Handler = PlainText
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

func New(tags ...Handler) *BBCode {
	return NewWithConfig(defaultConfig, tags...)
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
		w:      w,
		p:      b.tks.getParser(t),
		bbCode: b,
	}
	p.Process("")
	return p.err
}

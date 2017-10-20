package bbcode

import (
	"bytes"
	"text/template"
)

const (
	// HTMLText is a text processor that will HTML encode all text output
	HTMLText htmlText = 0
	// PlainText is a text processor that just outputs all text with no change
	PlainText plainText = 0
)

type htmlText int

var (
	newLine     = []byte{'\n'}
	newLineHTML = []byte{'<', 'b', 'r', ' ', '/', '>'}
)

func (htmlText) Name() string {
	return ""
}

func (htmlText) Handle(p *Processor, text string) {
	for n, line := range bytes.Split([]byte(text), newLine) {
		if n > 0 {
			p.Write(newLineHTML)
		}
		template.HTMLEscape(p, line)
	}
}

type plainText int

func (plainText) Name() string {
	return ""
}

func (plainText) Handle(p *Processor, text string) {
	p.Write([]byte(text))
}

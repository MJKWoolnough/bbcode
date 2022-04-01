package bbcode

import (
	"bytes"
	"io"
	"strings"

	"vimagination.zapto.org/parser"
)

// Processor contains methods necessary for creating custom Handler's
type Processor struct {
	w      io.Writer
	err    error
	p      parser.Parser
	bbCode *BBCode
}

// Write writes to the underlying writer.
// The error is stored and does not need to be handled
func (p *Processor) Write(b []byte) (int, error) {
	if p.err != nil {
		return 0, p.err
	}
	var n int
	n, p.err = p.w.Write(b)
	return n, p.err
}

// Process will continue processing the bbCode until it gets to an end tag
// which matches the tag name given, or until it reaches the end of the input.
// It returns true if the end tag was found, or false otherwise.
func (p *Processor) Process(untilTag string) bool {
	for {
		switch t := p.Get().(type) {
		case Text:
			p.printText(t)
		case OpenTag:
			p.ProcessTag(t)
		case CloseTag:
			if strings.EqualFold(t.Name, untilTag) {
				return true
			}
			p.printCloseTag(t)
		default:
			return false
		}
	}
}

// GetContents grabs the raw contents of a tag and returns it as a string
func (p *Processor) GetContents(untilTag string) string {
	if p.err != nil {
		return ""
	}
	w := p.w
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	p.w = b
	t := p.bbCode.text
	p.bbCode.text = PlainText
Loop:
	for {
		switch t := p.Get().(type) {
		case Text:
			p.printText(t)
		case OpenTag:
			p.printOpenTag(t)
		case CloseTag:
			if strings.EqualFold(t.Name, untilTag) {
				break Loop
			}
			p.printCloseTag(t)
		default:
			break Loop
		}
	}
	p.bbCode.text = t
	p.w = w
	return string(b.Bytes())
}

// ProcessTag will process the given tag as normal
func (p *Processor) ProcessTag(t OpenTag) {
	h := p.getTagHandler(t.Name)
	if h == nil {
		p.printOpenTag(t)
	} else {
		var attr string
		if t.Attr != nil {
			attr = *t.Attr
		}
		h.Handle(p, attr)
	}
}

func (p *Processor) getTagHandler(name string) Handler {
	for _, tag := range p.bbCode.tags {
		if strings.EqualFold(tag.Name(), name) {
			return tag
		}
	}
	return nil
}

// Get returns the next token.
// It will be either a Text, OpenTag or a CloseTag
func (p *Processor) Get() interface{} {
	phrase, _ := p.p.GetPhrase()
	switch phrase.Type {
	case phraseText:
		text := make(Text, 0, len(phrase.Data))
		for _, t := range phrase.Data {
			text = append(text, t.Data)
		}
		return text
	case phraseOpen:
		tag := OpenTag{
			Name: phrase.Data[0].Data,
		}
		if len(phrase.Data) > 1 {
			tag.Attr = &phrase.Data[1].Data
		}
		return tag
	case phraseClose:
		return CloseTag{Name: phrase.Data[0].Data}
	}
	return nil
}

// Print writes the textual representation of the given token to the output,
// using the text Handler
func (p *Processor) Print(t interface{}) {
	switch t := t.(type) {
	case string:
		p.bbCode.text.Handle(p, t)
	case Text:
		p.printText(t)
	case OpenTag:
		p.printOpenTag(t)
	case CloseTag:
		p.printCloseTag(t)
	}
}

func (p *Processor) printText(t Text) {
	for _, str := range t {
		p.bbCode.text.Handle(p, str)
	}
}

func (p *Processor) printOpenTag(t OpenTag) {
	p.bbCode.text.Handle(p, p.bbCode.tks.openTag)
	p.bbCode.text.Handle(p, t.Name)
	if t.Attr != nil {
		p.bbCode.text.Handle(p, p.bbCode.tks.attributeSep)
		p.bbCode.text.Handle(p, *t.Attr)
	}
	p.bbCode.text.Handle(p, p.bbCode.tks.closeTag)
}

func (p *Processor) printCloseTag(t CloseTag) {
	p.bbCode.text.Handle(p, p.bbCode.tks.openTag)
	p.bbCode.text.Handle(p, p.bbCode.tks.closingTag)
	p.bbCode.text.Handle(p, t.Name)
	p.bbCode.text.Handle(p, p.bbCode.tks.closeTag)
}

// Text is a token containing simple textual data
type Text []string

// OpenTag is a token containing the name of the tag and a possible attribute.
type OpenTag struct {
	Name string
	Attr *string
}

// CloseTag is a token containing the name of the tag
type CloseTag struct {
	Name string
}

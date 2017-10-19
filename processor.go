package bbcode

import (
	"io"
	"strings"

	"github.com/MJKWoolnough/parser"
)

type Processor struct {
	w    io.Writer
	err  error
	p    parser.Parser
	text Handler
	tags []Handler
}

func (p *Processor) Write(b []byte) (int, error) {
	if p.err != nil {
		return 0, p.err
	}
	var n int
	n, p.err = p.w.Write(b)
	return n, p.err
}

func (p *Processor) Process(untilTag string) bool {
	for {
		switch t := p.Get().(type) {
		case Text:
			for _, s := range t {
				p.text.Handle(p, s)
			}
		case OpenTag:
			p.ProcessTag(t)
		case CloseTag:
			if strings.ToLower(t.Name) == untilTag {
				return true
			}
			p.printCloseTag(t)
		default:
			return false
		}
	}
}

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
	for _, tag := range p.tags {
		if tag.Name() == name {
			return tag
		}
	}
	return nil
}

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

func (p *Processor) Print(t interface{}) {
	switch t := t.(type) {
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
		p.text.Handle(p, str)
	}
}

func (p *Processor) printOpenTag(t OpenTag) {
	p.text.Handle(p, tagBytes[:1])
	p.text.Handle(p, t.Name)
	if t.Attr != nil {
		p.text.Handle(p, tagBytes[2:3])
		p.text.Handle(p, *t.Attr)
	}
	p.text.Handle(p, tagBytes[3:])
}

const tagBytes = "[/=]"

func (p *Processor) printCloseTag(t CloseTag) {
	p.text.Handle(p, tagBytes[:2])
	p.text.Handle(p, t.Name)
	p.text.Handle(p, tagBytes[3:])
}

type Text []string

type OpenTag struct {
	Name string
	Attr *string
}

type CloseTag struct {
	Name string
}

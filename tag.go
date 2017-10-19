package bbcode

import "strings"

type Handler interface {
	Name() string
	Handle(*Processor, string)
}

type simpleTag struct {
	name        string
	open, close []byte
}

func Tag(name string, open, close []byte) Handler {
	return &simpleTag{
		name:  strings.ToLower(name),
		open:  open,
		close: close,
	}
}

func (s *simpleTag) Name() string {
	return s.name
}

func (s *simpleTag) Open(p *Processor, _ string) {
	p.Write(s.open)
}

func (s *simpleTag) Close(p *Processor) {
	p.Write(s.close)
}

func (s *simpleTag) Handle(p *Processor, a string) {
	s.Open(p, a)
	p.Process(s.name)
	s.Close(p)
}

type attributeTag struct {
	name                   string
	open, openClose, close []byte
	filter                 func(string) []byte
}

func AttributeTag(name string, open, openClose, close []byte, filter func(string) []byte) Handler {
	return &attributeTag{
		name:      strings.ToLower(name),
		open:      open,
		openClose: openClose,
		close:     close,
		filter:    filter,
	}
}

func (a *attributeTag) Name() string {
	return a.name
}

func (a *attributeTag) Open(p *Processor, attr string) {
	p.Write(a.open)
	p.Write(a.filter(attr))
	p.Write(a.openClose)
}

func (a *attributeTag) Close(p *Processor) {
	p.Write(a.close)
}

func (a *attributeTag) Handle(p *Processor, attr string) {
	a.Open(p, attr)
	p.Process(a.name)
	a.Close(p)
}

type OpenClose interface {
	Name() string
	Open(*Processor, string)
	Close(*Processor)
}

type filterTag struct {
	OpenClose
	filter func(string) bool
}

func FilterTag(o OpenClose, filter func(string) bool) Handler {
	return &filterTag{
		OpenClose: o,
		filter:    filter,
	}
}

func (f *filterTag) Handle(p *Processor, attr string) {
	f.Open(p, attr)
	allowText := f.filter("")
	name := f.Name()
Loop:
	for {
		switch t := p.Get().(type) {
		case Text:
			if allowText {
				p.Print(t)
			}
		case OpenTag:
			if f.filter(t.Name) {
				p.ProcessTag(t)
			} else if allowText {
				p.Print(t)
			}
		case CloseTag:
			if t.Name == name {
				break Loop
			}
			if allowText {
				p.Print(p)
			}
		default:
			break Loop
		}
	}
	f.Close(p)
}

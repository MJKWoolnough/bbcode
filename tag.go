package bbcode

type Handler interface {
	Name() string
	Handle(*Processor, string)
}

type SimpleTag struct {
	TagName           string
	TagOpen, TagClose []byte
}

func (s *SimpleTag) Name() string {
	return s.TagName
}

func (s *SimpleTag) Open(p *Processor, _ string) {
	p.Write(s.TagOpen)
}

func (s *SimpleTag) Close(p *Processor) {
	p.Write(s.TagClose)
}

func (s *SimpleTag) Handle(p *Processor, a string) {
	s.Open(p, a)
	p.Process(s.TagName)
	s.Close(p)
}

type AttributeTag struct {
	TagName                         string
	TagOpen, TagOpenClose, TagClose []byte
	Filter                          func(string) []byte
}

func (a *AttributeTag) Name() string {
	return a.TagName
}

func (a *AttributeTag) Open(p *Processor, attr string) {
	p.Write(a.TagOpen)
	p.Write(a.Filter(attr))
	p.Write(a.TagOpenClose)
}

func (a *AttributeTag) Close(p *Processor) {
	p.Write(a.TagClose)
}

func (a *AttributeTag) Handle(p *Processor, attr string) {
	a.Open(p, attr)
	p.Process(a.TagName)
	a.Close(p)
}

type OpenClose interface {
	Name() string
	Open(*Processor, string)
	Close(*Processor)
}

type FilterTag struct {
	OpenClose
	Filter func(string) bool
}

func (f *FilterTag) Handle(p *Processor, attr string) {
	f.Open(p, attr)
	allowText := f.Filter("")
	name := f.Name()
Loop:
	for {
		switch t := p.Get().(type) {
		case Text:
			if allowText {
				p.Print(t)
			}
		case OpenTag:
			if f.Filter(t.Name) {
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

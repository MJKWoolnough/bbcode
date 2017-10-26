package bbcodehtml

import (
	"strings"
	"unicode"

	"github.com/MJKWoolnough/bbcode"
)

type list struct{}

func (list) Name() string {
	return "list"
}

var (
	uListOpen  = []byte("<ul>")
	uListClose = []byte("</ul>")
	liOpen     = []byte("<li>")
	liClose    = []byte("</li>")
	oaList     = []byte("<ol type=\"a\">")
	oAList     = []byte("<ol type=\"A\">")
	oiList     = []byte("<ol type=\"i\">")
	oIList     = []byte("<ol type=\"I\">")
	o1List     = []byte("<ol type=\"1\">")
	oListClose = []byte("</ol>")
)

func (list) Handle(p *bbcode.Processor, attr string) {
	closer := oListClose
	switch attr {
	case "a":
		p.Write(oaList)
	case "A":
		p.Write(oAList)
	case "i":
		p.Write(oiList)
	case "I":
		p.Write(oIList)
	case "1":
		p.Write(o1List)
	default:
		p.Write(uListOpen)
		closer = uListClose
	}

	t := p.Get()

	var processed bool

	if text, ok := t.(bbcode.Text); ok {
		if strings.HasPrefix(strings.TrimLeftFunc(text[0], unicode.IsSpace), "*") {
			p.Write(liOpen)
			handleLiText(p, text)
			p.Write(liClose)
			processed = true
		}
	}

	if !processed {
	Loop:
		for {
			switch t := t.(type) {
			case bbcode.Text:
			case bbcode.OpenTag:
				if t.Name == "*" {
					if handleLi(p) {
						break Loop
					}
				}
			case bbcode.CloseTag:
				if t.Name == "list" {
					break Loop
				}
			default:
				break Loop
			}
			t = p.Get()
		}
	}

	p.Write(closer)
}

func handleLi(p *bbcode.Processor) bool {
	p.Write(liOpen)
	for {
		switch t := p.Get().(type) {
		case bbcode.Text:
			p.Print(t)
		case bbcode.OpenTag:
			if t.Name == "*" {
				p.Write(liClose)
				p.Write(liOpen)
			} else {
				p.ProcessTag(t)
			}
		case bbcode.CloseTag:
			switch strings.ToLower(t.Name) {
			case "*":
				p.Write(liClose)
				return false
			case "list":
				p.Write(liClose)
				return true
			default:
				p.Print(t)
			}
		}
	}
}

func handleLiText(p *bbcode.Processor, s interface{}) {
	var (
		start       = true
		lastNewLine [2]int
	)
Loop:
	for {
		switch t := s.(type) {
		case bbcode.Text:
			for m, text := range t {
				for n, c := range text {
					if c == '\n' {
						lastNewLine[0] = m
						lastNewLine[1] = n
						start = true
					} else if start {
						if c == '*' {
							p.Print(t[:lastNewLine[0]])
							t = t[lastNewLine[0]:]
							p.Print(t[0][:lastNewLine[1]])
							t[0] = t[0][:lastNewLine[1]+1]
							s = t
							start = false
							continue Loop
						} else if !unicode.IsSpace(c) {
							start = false
						}
					}
				}
			}
			p.Print(t)
		case bbcode.OpenTag:
			p.ProcessTag(t)
		case bbcode.CloseTag:
			if strings.ToLower(t.Name) == "list" {
				return
			}
			p.Print(t)
		default:
			return
		}
		start = false
		s = p.Get()
	}
}

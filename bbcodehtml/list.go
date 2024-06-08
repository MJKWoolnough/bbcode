package bbcodehtml

import (
	"strings"
	"unicode"

	"vimagination.zapto.org/bbcode"
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

	if text, ok := t.(bbcode.Text); ok && strings.HasPrefix(strings.TrimLeftFunc(text[0], unicode.IsSpace), "*") {
		p.Write(liOpen)
		handleLiText(p, text)
		p.Write(liClose)

		processed = true
	}

	if !processed {
	Loop:
		for {
			switch t := t.(type) {
			case bbcode.Text:
			case bbcode.OpenTag:
				if t.Name == "*" && handleLi(p) {
					break Loop
				}
			case bbcode.CloseTag:
				if strings.EqualFold(t.Name, "list") {
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
			if t.Name == "*" {
				p.Write(liClose)

				return false
			} else if strings.EqualFold(t.Name, "list") {
				p.Write(liClose)

				return true
			}

			p.Print(t)
		default:
			p.Write(liClose)

			return true
		}
	}
}

func handleLiText(p *bbcode.Processor, s interface{}) {
	var (
		start       = true
		lastNewLine [2]int
		firstDone   bool
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
							if firstDone {
								p.Print(t[:lastNewLine[0]])
								p.Print(text[:lastNewLine[1]])
								p.Write(liClose)
								p.Write(liOpen)
							} else {
								firstDone = true
							}

							t = t[lastNewLine[0]:]
							t[0] = t[0][n+1:]
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
			if strings.EqualFold(t.Name, "list") {
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

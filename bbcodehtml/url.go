package bbcodehtml

import (
	nurl "net/url"

	"github.com/MJKWoolnough/bbcode"
)

var (
	urlOpen  = []byte("<a href=\"")
	urlClose = []byte("</a>")
)

type url struct{}

func (url) Name() string {
	return "url"
}

func (url) Handle(p *bbcode.Processor, attr string) {
	if attr != "" {
		u, err := nurl.Parse(attr)
		if err == nil {
			p.Write(urlOpen)
			p.Write([]byte(u.String()))
			p.Write(attrTagClose)
			p.Process("url")
			p.Write(urlClose)
		} else {
			p.Process("url")
		}
	} else {
		utxt := p.GetContents("url")
		u, err := nurl.Parse(attr)
		if err == nil {
			p.Write(urlOpen)
			p.Write([]byte(u.String()))
			p.Write(attrTagClose)
			p.Print(utxt)
			p.Write(urlClose)
		} else {
			p.Print(utxt)
		}
	}
}

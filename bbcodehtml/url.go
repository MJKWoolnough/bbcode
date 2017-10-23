package bbcodehtml

import (
	"net/url"

	"github.com/MJKWoolnough/bbcode"
)

var (
	urlOpen  = []byte("<a href=\"")
	urlClose = []byte("</a>")
)

type urlT struct{}

func (urlT) Name() string {
	return "url"
}

func (urlT) Handle(p *bbcode.Processor, attr string) {
	if attr != "" {
		u, err := url.Parse(attr)
		if err == nil {
			p.Write(urlOpen)
			p.Write([]byte(u.String()))
			p.Write(attrClose)
			p.Write(tagClose)
			p.Process("url")
			p.Write(urlClose)
		} else {
			p.Process("url")
		}
	} else {
		utxt := p.GetContents("url")
		u, err := url.Parse(attr)
		if err == nil {
			p.Write(urlOpen)
			p.Write([]byte(u.String()))
			p.Write(attrClose)
			p.Write(tagClose)
			p.Print(utxt)
			p.Write(urlClose)
		} else {
			p.Print(utxt)
		}
	}
}

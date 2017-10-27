package bbcodehtml

import (
	"fmt"
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
			fmt.Fprintf(p, "<a href=%q>", u)
			p.Process("url")
			p.Write(urlClose)
		} else {
			p.Process("url")
		}
	} else {
		attr = p.GetContents("url")
		u, err := nurl.Parse(attr)
		if err == nil {
			fmt.Fprintf(p, "<a href=%q>", u)
			p.Print(attr)
			p.Write(urlClose)
		} else {
			p.Print(attr)
		}
	}
}

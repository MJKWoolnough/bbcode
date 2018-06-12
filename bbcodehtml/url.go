package bbcodehtml

import (
	"fmt"
	nurl "net/url"

	"vimagination.zapto.org/bbcode"
)

const (
	urlTag  = "url"
	urlOpen = "<a href=%q>"
)

var urlClose = []byte("</a>")

type url struct{}

func (url) Name() string {
	return "url"
}

func (url) Handle(p *bbcode.Processor, attr string) {
	if attr != "" {
		u, err := nurl.Parse(attr)
		if err == nil {
			fmt.Fprintf(p, urlOpen, u)
			p.Process(urlTag)
			p.Write(urlClose)
		} else {
			p.Process(urlTag)
		}
	} else {
		attr = p.GetContents(urlTag)
		u, err := nurl.Parse(attr)
		if err == nil {
			fmt.Fprintf(p, urlOpen, u)
			p.Print(attr)
			p.Write(urlClose)
		} else {
			p.Print(attr)
		}
	}
}

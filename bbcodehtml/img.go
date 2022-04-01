package bbcodehtml

import (
	"bytes"
	"fmt"
	"html/template"
	nurl "net/url"
	"strings"

	"vimagination.zapto.org/bbcode"
)

var (
	imgOpen = []byte("<img")
	altOpen = []byte(" alt=\"")
)

type img struct{}

func (img) Name() string {
	return "img"
}

func (img) Handle(p *bbcode.Processor, attr string) {
	if u, err := nurl.Parse(p.GetContents("img")); err == nil {
		switch strings.ToLower(u.Scheme) {
		case "http", "https":
			p.Write(imgOpen)
			if attr != "" {
				p.Write(altOpen)
				var b bytes.Buffer
				template.HTMLEscape(&b, []byte(attr))
				p.Write(b.Bytes())
				p.Write(attrClose)
			}
			fmt.Fprintf(p, " src=%q />", u)
		}
	}
}

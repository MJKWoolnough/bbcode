package bbcodehtml

import (
	"fmt"
	"html/template"
	nurl "net/url"
	"strings"

	"github.com/MJKWoolnough/bbcode"
	"github.com/MJKWoolnough/memio"
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
				var b memio.Buffer
				template.HTMLEscape(&b, []byte(attr))
				p.Write(b)
				p.Write(attrClose)
			}
			fmt.Fprintf(p, " src=%q />", u)
		}
	}
}

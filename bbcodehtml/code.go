package bbcodehtml

import "vimagination.zapto.org/bbcode"

type code struct{}

func (code) Name() string {
	return "code"
}

var (
	codeOpen  = []byte("<pre>")
	codeClose = []byte("</pre>")
)

func (code) Handle(p *bbcode.Processor, _ string) {
	p.Write(codeOpen)
	p.Print(p.GetContents("code"))
	p.Write(codeClose)
}

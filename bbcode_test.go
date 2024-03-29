package bbcode

import (
	"bytes"
	"testing"
)

var (
	_ Handler   = &Tag{}
	_ Handler   = &AttributeTag{}
	_ OpenClose = &Tag{}
	_ OpenClose = &AttributeTag{}
)

func passFilter(a string) []byte {
	return []byte(a)
}

func TestParse(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	b := New(
		HTMLText,
		NewTag("b", []byte("<b>"), []byte("</b>")),
		NewTag("i", []byte("<i>"), []byte("</i>")),
		NewAttributeTag(
			"tester",
			[]byte("<span"),
			[]byte(">"),
			[]byte(" style=\"color: "),
			[]byte("\""),
			[]byte("</span>"),
			AttrFilterFunc(passFilter),
		),
	)
	for n, test := range []struct {
		Input, Output string
	}{
		{
			"[B]Bolded ] [i]Bolded-Italic[/I] [ [/b] [I]Just Italic[/I]",
			"<b>Bolded ] <i>Bolded-Italic</i> [ </b> <i>Just Italic</i>",
		},
		{
			"[b][UnknownTag=AttributeFoo][/UnknownTAG]",
			"<b>[UnknownTag=AttributeFoo][/UnknownTAG]</b>",
		},
		{
			"[tester=#fff]White",
			"<span style=\"color: #fff\">White</span>",
		},
	} {
		buf.Reset()
		b.ConvertString(buf, test.Input)
		if str := buf.String(); str != test.Output {
			t.Errorf("test %d: expecting %q, got %q", n+1, test.Output, str)
		}
	}
}

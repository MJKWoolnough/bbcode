package bbcodehtml

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/MJKWoolnough/bbcode"
	"github.com/MJKWoolnough/memio"
)

var (
	alignLeft   = []byte("left")
	alignRight  = []byte("right")
	alignCentre = []byte("center")
	alignFull   = []byte("justify")

	leftOpen   = []byte("<div style=\"text-align: left\">")
	centreOpen = []byte("<div style=\"text-align: center\">")
	rightOpen  = []byte("<div style=\"text-align: right\">")
	fullOpen   = []byte("<div style=\"text-align: justify\">")

	divOpen        = []byte("<div>")
	divPartialOpen = []byte("<div")
	divClose       = []byte("</div>")

	spanOpen        = []byte("<span>")
	spanPartialOpen = []byte("<span")
	spanClose       = []byte("</span>")

	tagClose  = []byte(">")
	attrClose = []byte("\"")
	sizeClose = []byte("pt\"")

	alignAttr  = []byte(" style=\"text-align: ")
	colourAttr = []byte(" style=\"color: ")
	fontAttr   = []byte(" style=\"font-family: ")
	sizeAttr   = []byte(" style=\"font-size: ")
)

var (
	Align = bbcode.NewAttributeTag("align", divPartialOpen, tagClose, alignAttr, attrClose, divClose, alignFilter)

	LeftAlign   = bbcode.NewTag("left", leftOpen, divClose)
	CentreAlign = bbcode.NewTag("centre", centreOpen, divClose)
	CenterAlign = bbcode.NewTag("center", centreOpen, divClose)
	RightAlign  = bbcode.NewTag("right", rightOpen, divClose)
	FullAlign   = bbcode.NewTag("full", fullOpen, divClose)

	Color  = bbcode.NewAttributeTag("color", spanPartialOpen, tagClose, colourAttr, attrClose, spanClose, Colours.Filter)
	Colour = bbcode.NewAttributeTag("colour", spanPartialOpen, tagClose, colourAttr, attrClose, spanClose, Colours.Filter)

	Font = bbcode.NewAttributeTag("font", spanPartialOpen, tagClose, fontAttr, attrClose, spanClose, Fonts.Filter)

	Bold         = bbcode.NewTag("b", []byte("<b>"), []byte("</b>"))
	Italic       = bbcode.NewTag("i", []byte("<i>"), []byte("</i>"))
	Strikethough = bbcode.NewTag("s", []byte("<s>"), []byte("</s>"))
	Underline    = bbcode.NewTag("u", []byte("<u>"), []byte("</u>"))

	Size = bbcode.NewAttributeTag("size", spanPartialOpen, tagClose, sizeAttr, sizeClose, spanClose, sizeFilter)

	Heading1 = bbcode.NewTag("h1", []byte("<h1>"), []byte("</h1>"))
	Heading2 = bbcode.NewTag("h2", []byte("<h2>"), []byte("</h2>"))
	Heading3 = bbcode.NewTag("h3", []byte("<h3>"), []byte("</h3>"))
	Heading4 = bbcode.NewTag("h4", []byte("<h4>"), []byte("</h4>"))
	Heading5 = bbcode.NewTag("h5", []byte("<h5>"), []byte("</h5>"))
	Heading6 = bbcode.NewTag("h6", []byte("<h6>"), []byte("</h6>"))
	Heading7 = bbcode.NewTag("h7", []byte("<h7>"), []byte("</h7>"))

	Quote = bbcode.AttributeTag("quote", []byte("<fieldset class=\"quote\">"), []byte("<blockquote>"), []byte("<legend>"), []byte("</legend>"), []byte("</blockquote></fieldset>"), notEmptyFilter)

	Code code
	// Image img  // TODO
	// List  list // TODO
	Table table
	// URL   url  // TODO
)

func alignFilter(attr string) []byte {
	switch strings.ToLower(attr) {
	case "left":
		return alignLeft
	case "center", "centre":
		return alignCentre
	case "right":
		return alignRight
	case "full", "justify", "fulljustify", "full-justify", "full justify":
		return alignFull
	}
	return nil
}

func sizeFilter(attr string) []byte {
	n, _ := strconv.Atoi(attr)
	if n >= 1 && n <= 50 {
		return []byte(attr)
	}
	return nil
}

func notEmptyFilter(attr string) []byte {
	if len(attr) == 0 {
		return nil
	} else if !strings.ContainsAny(attr, "'\"&<>\000") {
		return []byte(attr)
	}
	var b memio.Buffer
	template.HTMLEscape(&b, []byte(attr))
	return b
}
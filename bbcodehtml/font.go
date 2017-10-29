package bbcodehtml

import "strings"

// Fonts is a map of font names and their css expansion for use as an
// attribute filter for bbcode.AttributeTag
var Fonts = fonts{
	"georgia":         []byte("Georgia, serif"),
	"times":           []byte("\\\"Times New Roman\\\", Times, serif"),
	"times new roman": []byte("\\\"Times New Roman\\\", Times, serif"),
	"arial":           []byte("Arial, Helvetica, sans-serif"),
	"arial black":     []byte("\\\"Arial Black\\\", Gadget, sans-serif"),
	"comic sans ms":   []byte("\\\"Comic Sans MS\\\", cursive, sans-serif"),
	"comic sans":      []byte("\\\"Comic Sans MS\\\", cursive, sans-serif"),
	"impact":          []byte("Impact, Charcoal, sans-serif"),
	"verdana":         []byte("Verdana, Geneva, sans-serif"),
	"courier":         []byte("\\\"Courier New\\\", Courier, monospace"),
	"lucida console":  []byte("\\\"Lucida Console\\\", Monaco, monospace"),
	"serif":           []byte("serif"),
	"sans serif":      []byte("sans-serif"),
	"monospace":       []byte("monospace"),
}

type fonts map[string][]byte

func (f fonts) AttrFilter(attr string) []byte {
	if font, ok := f[strings.ToLower(attr)]; ok {
		return font
	}
	return nil
}

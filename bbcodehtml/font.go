package bbcodehtml

import (
	"strings"
)

var (
	georgia    = []byte("Georgia, serif")
	times      = []byte("\\\"Times New Roman\\\", Times, serif")
	arial      = []byte("Arial, Helvetica, sans-serif")
	arialBlack = []byte("\\\"Arial Black\\\", Gadget, sans-serif")
	comicSans  = []byte("\\\"Comic Sans MS\\\", cursive, sans-serif")
	impact     = []byte("Impact, Charcoal, sans-serif")
	verdana    = []byte("Verdana, Geneva, sans-serif")
	courier    = []byte("\\\"Courier New\\\", Courier, monospace")
	lucida     = []byte("\\\"Lucida Console\\\", Monaco, monospace")
	sansSerif  = impact[18:]
	serif      = sansSerif[5:]
	monospace  = courier[26:]
)

// Fonts is a map of font names and their CSS expansion for use as an
// attribute filter for bbcode.AttributeTag.
var Fonts = fonts{
	"georgia":         georgia,
	"times":           times,
	"times new roman": times,
	"arial":           arial,
	"arial black":     arialBlack,
	"comic sans ms":   comicSans,
	"comic sans":      comicSans,
	"impact":          impact,
	"verdana":         verdana,
	"courier":         courier,
	"lucida console":  lucida,
	"serif":           serif,
	"sans serif":      sansSerif,
	"monospace":       monospace,
}

type fonts map[string][]byte

func (f fonts) AttrFilter(attr string) []byte {
	if font, ok := f[strings.ToLower(attr)]; ok {
		return font
	}

	return nil
}

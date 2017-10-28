package bbcodehtml

import (
	"testing"
)

func TestColour(t *testing.T) {
	testAttr(t, []attrinout{
		{"RED", []byte("#FF0000")},
		{"Red", []byte("#FF0000")},
		{"red", []byte("#FF0000")},
		{"rED", []byte("#FF0000")},
		{"blue", []byte("#0000FF")},
		{"#f00", []byte("#f00")},
		{"#F00", []byte("#F00")},
		{"#FA1258", []byte("#FA1258")},
		{"ffffff", []byte("#ffffff")},
		{"000", []byte("#000")},
		{"Not-A-Colour", nil},
	})
}

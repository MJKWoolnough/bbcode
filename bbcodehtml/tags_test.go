package bbcodehtml

import (
	"bytes"
	"testing"

	"vimagination.zapto.org/bbcode"
	"vimagination.zapto.org/memio"
)

type inout struct {
	Input, Output string
}

func testTag(t *testing.T, tests []inout, types ...bbcode.Handler) {
	t.Parallel()
	b := bbcode.New(types...)
	var buf memio.Buffer
	for n, test := range tests {
		buf = buf[:0]
		if err := b.ConvertString(&buf, test.Input); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
			continue
		}
		if res := string(buf); res != test.Output {
			t.Errorf("test %d: expecting %q, got %q", n+1, test.Output, res)
		}
	}
}

type attrinout struct {
	Input  string
	Output []byte
}

func testAttr(t *testing.T, tests []attrinout) {
	t.Parallel()
	for n, test := range tests {
		if output := Colours.AttrFilter(test.Input); !bytes.Equal(test.Output, output) {
			t.Errorf("test %d: expecting %s, got %s", n+1, test.Output, output)
		}
	}
}

package bbcodehtml

import (
	"testing"
)

func TestCode(t *testing.T) {
	testTag(t, []inout{
		{"[code][/code]", "<pre></pre>"},
		{"[code][b][/code]", "<pre>[b]</pre>"},
		{"[code][b]Hello<[/b][/code]", "<pre>[b]Hello&lt;[/b]</pre>"},
		{"[code][code][/code]", "<pre>[code]</pre>"},
	}, Code, Bold)
}

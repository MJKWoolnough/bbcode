package bbcodehtml

import (
	"testing"
)

func TestList(t *testing.T) {
	testTag(t, []inout{
		{"[list][*][b]1[/b][/*][/list]", "<ul><li><b>1</b></li></ul>"},
		{"[list][*][b]1[/b][/*][*]2[/*][/list]", "<ul><li><b>1</b></li><li>2</li></ul>"},
		{"[list][*][b]1[/b][*]2[/list]", "<ul><li><b>1</b></li><li>2</li></ul>"},
		{"[list][*][b]1[/b][*]2[/*][/list]", "<ul><li><b>1</b></li><li>2</li></ul>"},
		{"[list=1][*]1[*]2", "<ol type=\"1\"><li>1</li><li>2</li></ol>"},
		{"[list=a][*]1[*]2", "<ol type=\"a\"><li>1</li><li>2</li></ol>"},
		{"[list=A][*]1[*]2", "<ol type=\"A\"><li>1</li><li>2</li></ol>"},
		{"[list=i][*]1[*]2", "<ol type=\"i\"><li>1</li><li>2</li></ol>"},
		{"[list=I][*]1[*]2", "<ol type=\"I\"><li>1</li><li>2</li></ol>"},
		{"[list=F][*]1[*]2", "<ul><li>1</li><li>2</li></ul>"},
		{
			"[list]\n" +
				" * Hello\n" +
				"* Beep\n" +
				"	 Boop\n" +
				"	 *Foo",
			"<ul><li> Hello</li><li> Beep<br />	 Boop</li><li>Foo</li></ul>",
		},
		{
			"[list=a]\n" +
				" * Hello\n" +
				"* Beep\n" +
				"	 Boop\n" +
				"	 *Foo[/list]",
			"<ol type=\"a\"><li> Hello</li><li> Beep<br />	 Boop</li><li>Foo</li></ol>",
		},
	}, List, Bold)
}

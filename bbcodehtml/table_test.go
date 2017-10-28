package bbcodehtml

import (
	"testing"
)

func TestTable(t *testing.T) {
	testTag(t, []inout{
		{"[table][/table]", "<table></table>"},
		{"[table][b]a[/b][/table]", "<table></table>"},
		{"[table][tr][b]a[/b][/tr][/table]", "<table><tr></tr></table>"},
		{"[table][tr][td][b]a[/b][/td][/tr][/table]", "<table><tr><td><b>a</b></td></tr></table>"},
		{
			"[table]\n" +
				"	[thead]\n" +
				"		[tr]" +
				"			[th]Hello[/th]\n" +
				"		[/tr]\n" +
				"		[tr]\n" +
				"			[td]World[/td]\n" +
				"		[/tr]\n" +
				"	[/thead]\n" +
				"	[tfoot]\n" +
				"		[row]\n" +
				"			[col]Footer[/col]\n" +
				"		[/row]\n" +
				"	[/tfoot]\n" +
				"	[tr]\n" +
				"		[td][b]a[/b][/td]\n" +
				"	[/tr]\n" +
				"[/table]",
			"<table><thead><tr><th>Hello</th></tr><tr><td>World</td></tr></thead><tfoot><tr><td>Footer</td></tr></tfoot><tr><td><b>a</b></td></tr></table>",
		},
	}, Table, Bold)
}

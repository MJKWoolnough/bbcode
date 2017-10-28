package bbcodehtml

import (
	"testing"
)

func TestURL(t *testing.T) {
	testTag(t, []inout{
		{"[url=http://www.example.com]Example[/url]", "<a href=\"http://www.example.com\">Example</a>"},
		{"[url=http://www.example.com][b]E[/b]xample[/url]", "<a href=\"http://www.example.com\"><b>E</b>xample</a>"},
		{"[url]http://www.example.com[/url]", "<a href=\"http://www.example.com\">http://www.example.com</a>"},
		{"[url]http://www.example.com/[b]a[/b][/url]", "<a href=\"http://www.example.com/[b]a[/b]\">http://www.example.com/[b]a[/b]</a>"},
		{"[url=http://www.example.com/[b]a[/b]]T[b]e[/b]st[/url]", "<a href=\"http://www.example.com/[b\">a[/b]]T<b>e</b>st</a>"},
		{"[url]http://www.exa\"mple.com/a\"?b=\"1\"[/url]", "<a href=\"http://www.exa\\\"mple.com/a%22?b=\\\"1\\\"\">http://www.exa&#34;mple.com/a&#34;?b=&#34;1&#34;</a>"},
	}, URL, Bold)
}

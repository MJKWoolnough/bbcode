package bbcodehtml

import (
	"testing"
)

func TestImg(t *testing.T) {
	testTag(t, []inout{
		{"[img]http://www.example.com/image.png[/img]", "<img src=\"http://www.example.com/image.png\" />"},
		{"[img]http://www.example.com/image\".png[/img]", "<img src=\"http://www.example.com/image%22.png\" />"},
		{"[img]http://www.exam\"ple.com/image.png[/img]", "<img src=\"http://www.exam\\\"ple.com/image.png\" />"},
	}, Image, Bold)
}

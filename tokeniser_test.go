package bbcode

import (
	"testing"

	"github.com/MJKWoolnough/parser"
)

func TestTokeniser(t *testing.T) {
Loop:
	for n, test := range [...]struct {
		Input  string
		Output []parser.Phrase
	}{
		{},
		{
			Input: "A",
			Output: []parser.Phrase{
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "A",
						},
					},
				},
			},
		},
		{
			Input: "AB",
			Output: []parser.Phrase{
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "AB",
						},
					},
				},
			},
		},
		{
			Input: "[tag][/tag]",
			Output: []parser.Phrase{
				{
					Type: phraseOpen,
					Data: []parser.Token{
						{
							Type: tokenOpenTag,
							Data: "tag",
						},
					},
				},
				{
					Type: phraseClose,
					Data: []parser.Token{
						{
							Type: tokenCloseTag,
							Data: "tag",
						},
					},
				},
			},
		},
		{
			Input: "[tag=attr][/tag]",
			Output: []parser.Phrase{
				{
					Type: phraseOpen,
					Data: []parser.Token{
						{
							Type: tokenOpenTag,
							Data: "tag",
						},
						{
							Type: tokenTagAttribute,
							Data: "attr",
						},
					},
				},
				{
					Type: phraseClose,
					Data: []parser.Token{
						{
							Type: tokenCloseTag,
							Data: "tag",
						},
					},
				},
			},
		},
		{
			Input: "ABC[img]urlHere[/img]123[url=http://www.google.com]Link Here[/url]",
			Output: []parser.Phrase{
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "ABC",
						},
					},
				},
				{
					Type: phraseOpen,
					Data: []parser.Token{
						{
							Type: tokenOpenTag,
							Data: "img",
						},
					},
				},
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "urlHere",
						},
					},
				},
				{
					Type: phraseClose,
					Data: []parser.Token{
						{
							Type: tokenCloseTag,
							Data: "img",
						},
					},
				},
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "123",
						},
					},
				},
				{
					Type: phraseOpen,
					Data: []parser.Token{
						{
							Type: tokenOpenTag,
							Data: "url",
						},
						{
							Type: tokenTagAttribute,
							Data: "http://www.google.com",
						},
					},
				},
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "Link Here",
						},
					},
				},
				{
					Type: phraseClose,
					Data: []parser.Token{
						{
							Type: tokenCloseTag,
							Data: "url",
						},
					},
				},
			},
		},
		{
			Input: "PreText[NotATag[NowATag]",
			Output: []parser.Phrase{
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "PreText",
						},
						{
							Type: tokenText,
							Data: "[NotATag",
						},
					},
				},
				{
					Type: phraseOpen,
					Data: []parser.Token{
						{
							Type: tokenOpenTag,
							Data: "NowATag",
						},
					},
				},
			},
		},
		{
			Input: "PreText[NotATag=[NowATag]",
			Output: []parser.Phrase{
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "PreText",
						},
					},
				},
				{
					Type: phraseOpen,
					Data: []parser.Token{
						{
							Type: tokenOpenTag,
							Data: "NotATag",
						},
						{
							Type: tokenTagAttribute,
							Data: "[NowATag",
						},
					},
				},
			},
		},
		{
			Input: "PreText[NotATag=[StillNotATag",
			Output: []parser.Phrase{
				{
					Type: phraseText,
					Data: []parser.Token{
						{
							Type: tokenText,
							Data: "PreText",
						},
						{
							Type: tokenText,
							Data: "[NotATag",
						},
						{
							Type: tokenText,
							Data: "=[StillNotATag",
						},
					},
				},
			},
		},
	} {
		tk := newTokeniser(parser.NewStringTokeniser(test.Input))
		for m, phrase := range test.Output {
			p, err := tk.GetPhrase()
			if err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, err)
			}
			if phrase.Type != p.Type {
				t.Errorf("test %d.%d: expecting phrase type %d, got %d", n+1, m+1, phrase.Type, p.Type)
				break Loop
			}
			if len(phrase.Data) != len(p.Data) {
				t.Errorf("test %d.%d: expecting %d tokens, got %d", n+1, m+1, len(phrase.Data), len(p.Data))
				break Loop
			}
			for l, token := range phrase.Data {
				if token.Type != p.Data[l].Type {
					t.Errorf("test %d.%d.%d: expecting token type %d, got %d", n+1, m+1, l+1, token.Type, p.Data[l].Type)
					break Loop
				}
				if token.Data != p.Data[l].Data {
					t.Errorf("test %d.%d.%d: expecting token data %q, got %q", n+1, m+1, l+1, token.Data, p.Data[l].Data)
					break Loop
				}
			}
		}
		if p, _ := tk.GetPhrase(); p.Type != parser.PhraseDone {
			t.Errorf("test %d: more tokens to check", n+1)
		}
	}
}

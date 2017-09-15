package bbcode

import "github.com/MJKWoolnough/parser"

const (
	tokenText parser.TokenType = iota
	tokenOpenTag
	tokenTagAttribute
	tokenCloseTag
)

const (
	openTag      = "["
	closeTag     = "]"
	closingTag   = "/"
	validTagName = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*"
	attributeSep = "="
)

func newTokeniser(data string) *parser.Tokeniser {
	t := parser.NewStringTokeniser(data)
	t.TokeniserState(text)
	return &t
}

func text(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.ExceptRun(openTag)
	tk := parser.Token{
		tokenText,
		t.Get(),
	}
	if t.Peek() == -1 {
		return tk, done
	}
	return tk, opening
}

func opening(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept(openTag)
	if t.Peek() == rune(closingTag[0]) {
		return closing(t)
	}
	t.AcceptRun(validTagName)
	var (
		next parser.TokenFunc
		data string
	)
	switch t.Peek() {
	case rune(closeTag[0]):
		next = text
		data = t.Get()
		t.Accept(closeTag)
	case rune(attributeSep[0]):
		next = attribute
		data = t.Get()
		t.Accept(attributeSep)
	default:
		return text(t)
	}
	t.Get()
	data = data[1:]
	return parser.Token{
		tokenOpenTag,
		data,
	}, next
}

func closing(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept(closingTag)
	t.AcceptRun(validTagName)
	if t.Peek() == rune(closeTag[0]) {
		data := t.Get()
		t.Accept(closeTag)
		t.Get()
		return parser.Token{
			tokenCloseTag,
			data[2:],
		}, text
	}
	return text(t)
}

func attribute(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.ExceptRun(closeTag)
	if t.Peek() == -1 {
		return text(t)
	}
	data := t.Get()
	t.Accept(closeTag)
	t.Get()
	return parser.Token{
		tokenTagAttribute,
		data,
	}, text
}

func done(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return t.Done()
}

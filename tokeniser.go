package bbcode

import "github.com/MJKWoolnough/parser"

const (
	tokenText parser.TokenType = iota
	tokenOpenTag
	tokenTagAttribute
	tokenCloseTag
)

const (
	phraseText parser.PhraseType = iota
	phraseOpen
	phraseClose
)

const (
	openTag      = "["
	closeTag     = "]"
	closingTag   = "/"
	validTagName = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*"
	attributeSep = "="
)

func newTokeniser(t parser.Tokeniser) *parser.Parser {
	p := parser.New(t)
	p.TokeniserState(text)
	p.PhraserState(phraser)
	return &p
}

func text(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.ExceptRun(openTag)
	tk := parser.Token{
		tokenText,
		t.Get(),
	}
	if t.Peek() == -1 {
		return tk, (*parser.Tokeniser).Done
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

func phraser(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
	var phraseType parser.PhraseType
	if p.Accept(tokenText) {
		p.AcceptRun(tokenText)
		phraseType = phraseText
	} else if p.Accept(tokenOpenTag) {
		p.Accept(tokenTagAttribute)
		phraseType = phraseOpen
	} else if p.Accept(tokenCloseTag) {
		phraseType = phraseClose
	} else if p.Accept(parser.TokenDone) {
		return p.Done()
	} else if p.Accept(parser.TokenError) {
		return p.Error()
	}
	return parser.Phrase{
		Type: phraseType,
		Data: p.Get(),
	}, phraser
}

package bbcode

import (
	"errors"

	"github.com/MJKWoolnough/parser"
)

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

func newTokeniser(data string) *parser.Parser {
	p := parser.New(parser.NewStringTokeniser(data))
	p.TokeniserState(text)
	p.PhraserState(phraserText)
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

func phraserText(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
	var next parser.PhraseFunc
	switch p.AcceptRun(tokenText) {
	case parser.TokenDone:
		next = phraserDone
	case tokenOpenTag:
		next = phraserOpen
	case tokenCloseTag:
		next = phraserClose
	default:
		p.Err = errors.New("invalid state")
		return p.Error()
	}
	ts := p.Get()
	if len(ts) == 0 {
		return next()
	} else if len(ts) > 1 {
		var l int
		for _, t := range ts {
			l += string(t.Data)
		}
		str := make([]byte, 0, l)
		for _, t := range ts {
			str = append(str, t.Data...)
		}
		ts[0] = string(str)
		ts = ts[:1]
	}
	return parser.Phrase{
		phraseText,
		ts,
	}, next
}

func phraserOpen(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
	p.Accept(tokenOpenTag)
	p.Accept(tokenTagAttribute)
	return parser.Phrase{
		phraseOpen,
		p.Get(),
	}, phraserText
}

func phraserClose(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
	p.Accept(tokenCloseTag)
	return parser.Phrase{
		phraseOpen,
		p.Get(),
	}, phraserText
}

func phraserDone(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
	return p.Done()
}

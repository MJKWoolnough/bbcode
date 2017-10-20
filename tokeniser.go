package bbcode

import (
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

type tokeniser struct {
	openTag, closeTag, closingTag, validTagName, attributeSep string
	closeTagRune, closingTagRune, attributeSepRune            rune
}

func getTokeniser(c Config) tokeniser {
	return tokeniser{
		openTag:          string(c.TagOpen),
		closeTag:         string(c.TagClose),
		closingTag:       string(c.ClosingTag),
		validTagName:     string(c.Name),
		attributeSep:     string(c.AttributeSep),
		closeTagRune:     c.TagClose,
		closingTagRune:   c.ClosingTag,
		attributeSepRune: c.AttributeSep,
	}
}

func (tks *tokeniser) getParser(t parser.Tokeniser) parser.Parser {
	p := parser.New(t)
	p.TokeniserState(tks.text)
	p.PhraserState(tks.phraser)
	return p
}

func (tks *tokeniser) text(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.ExceptRun(tks.openTag)
	tk := parser.Token{
		tokenText,
		t.Get(),
	}
	if t.Peek() == -1 {
		if tk.Data == "" {
			return t.Done()
		}
		return tk, (*parser.Tokeniser).Done
	}
	if tk.Data == "" {
		return tks.opening(t)
	}
	return tk, tks.opening
}

func (tks *tokeniser) opening(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept(tks.openTag)
	if t.Peek() == tks.closingTagRune {
		return tks.closing(t)
	}
	if !t.Accept(tks.validTagName) {
		return tks.text(t)
	}
	t.AcceptRun(tks.validTagName)
	var (
		next parser.TokenFunc = tks.text
		data string
	)
	switch t.Peek() {
	case tks.closeTagRune:
		data = t.Get()
		t.Accept(tks.closeTag)
		t.Get()
	case tks.attributeSepRune:
		data = t.Get()
		t.Accept(tks.attributeSep)
		if t.ExceptRun(tks.closeTag) != tks.closeTagRune {
			return parser.Token{
				tokenText,
				data,
			}, tks.text
		}
		next = tks.attribute
	default:
		return tks.text(t)
	}
	data = data[1:]
	return parser.Token{
		tokenOpenTag,
		data,
	}, next
}

func (tks *tokeniser) closing(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept(tks.closingTag)
	if !t.Accept(tks.validTagName) {
		return tks.text(t)
	}
	t.AcceptRun(tks.validTagName)
	if t.Peek() == tks.closeTagRune {
		data := t.Get()
		t.Accept(tks.closeTag)
		t.Get()
		return parser.Token{
			tokenCloseTag,
			data[2:],
		}, tks.text
	}
	return tks.text(t)
}

func (tks *tokeniser) attribute(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	data := t.Get()
	t.Accept(tks.closeTag)
	t.Get()
	return parser.Token{
		tokenTagAttribute,
		data[1:],
	}, tks.text
}

func (tks *tokeniser) phraser(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
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
	}, tks.phraser
}

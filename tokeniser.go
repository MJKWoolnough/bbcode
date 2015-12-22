package bbcode

import (
	"io"

	"github.com/MJKWoolnough/parser"
)

type tokenType uint8

const (
	tokenText tokenType = iota
	tokenOpenTag
	tokenTagAttribute
	tokenCloseTag
	tokenDone
)

const (
	openTag      = "["
	closeTag     = "]"
	closingTag   = "/"
	validTagName = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*"
	attributeSep = "="
)

type token struct {
	typ  tokenType
	data string
}

type stateFn func() (token, stateFn)

type tokeniser struct {
	p     parser.Parser
	state stateFn
	err   error
}

func newTokeniser(data string) *tokeniser {
	t := &tokeniser{
		p: parser.NewStringParser(data),
	}
	t.state = t.text
	return t
}

func (t *tokeniser) GetToken() (token, error) {
	var tk token
	tk, t.state = t.state()
	return tk, t.err
}

func (t *tokeniser) text() (token, stateFn) {
	t.p.ExceptRun(openTag)
	tk := token{
		tokenText,
		t.p.Get(),
	}
	if t.p.Peek() == -1 {
		return tk, t.done
	}
	return tk, t.tag
}

func (t *tokeniser) tag() (token, stateFn) {
	t.p.Accept(openTag)
	if t.p.Peek() == rune(closingTag[0]) {
		return t.closingTag()
	}
	t.p.AcceptRun(validTagName)
	var next stateFn
	switch t.p.Peek() {
	case rune(closeTag[0]):
		next = t.text
	case rune(attributeSep[0]):
		next = t.attribute
	default:
		return t.text()
	}
	data := t.p.Get()
	t.p.Accept(closeTag + attributeSep)
	t.p.Get()
	data = data[1:]
	return token{
		tokenOpenTag,
		data,
	}, next
}

func (t *tokeniser) closingTag() (token, stateFn) {
	t.p.Accept(closingTag)
	t.p.AcceptRun(validTagName)
	if t.p.Peek() == rune(closeTag[0]) {
		data := t.p.Get()
		t.p.Accept(closeTag)
		t.p.Get()
		return token{
			tokenCloseTag,
			data[2:],
		}, t.text
	}
	return t.text()
}

func (t *tokeniser) attribute() (token, stateFn) {
	t.p.ExceptRun(closeTag)
	if t.p.Peek() == -1 {
		return t.text()
	}
	data := t.p.Get()
	t.p.Accept(closeTag)
	t.p.Get()
	return token{
		tokenTagAttribute,
		data,
	}, t.text
}

func (t *tokeniser) done() (token, stateFn) {
	t.err = io.EOF
	return token{tokenDone, ""}, t.done
}

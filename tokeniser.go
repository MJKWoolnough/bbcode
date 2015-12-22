package bbcode

import (
	"io"

	"github.com/MJKWoolnough/parser"
)

type tokenType uint8

const (
	tokenError tokenType = iota
	tokenText
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
	if t.err == io.EOF {
		return token{tokenDone, ""}, io.EOF
	}
	var tk token
	tk, t.state = t.state()
	if t.err == io.EOF {
		if tk.typ == tokenError {
			t.err = io.ErrUnexpectedEOF
		} else {
			return tk, nil
		}
	}
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
	typ := tokenOpenTag
	if t.p.Peek() == rune(closingTag[0]) {
		t.p.Accept(closingTag)
		typ = tokenCloseTag
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
	if typ == tokenCloseTag {
		data = data[1:]
	}
	return token{
		typ,
		data,
	}, next
}

func (t *tokeniser) attribute() (token, stateFn) {
	t.p.ExceptRun(closeTag)
	if t.p.Peek() == -1 {
		return t.text()
	}
	return token{
		tokenTagAttribute,
		t.p.Get(),
	}, t.text
}

func (t *tokeniser) done() (token, stateFn) {
	t.err = io.EOF
	return token{tokenDone, ""}, t.done
}

func (t *tokeniser) errorfn() (token, stateFn) {
	return token{
		tokenError,
		t.err.Error(),
	}, t.errorfn
}

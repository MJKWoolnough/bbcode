// Package bbcode implements a bbcode parser and converter
package bbcode

import (
	"io"

	"github.com/MJKWoolnough/parser"
)

// Converter is an interface for converting the tag tree into a string
type Converter interface {
	Open(*Tag) string
	Close(*Tag) string
}

// BBCodeConverter is a Converter that will translate a parsed tree back into
// BBCode
type BBCodeConverter struct{}

// Open creates the opening of a BBCode tag
func (BBCodeConverter) Open(t *Tag) string {
	if t.Attribute != "" {
		return "[" + t.Name + "=" + t.Attribute + "]"
	}
	return "[" + t.Name + "]"
}

// Close creates the closing of a BBCode tag
func (BBCodeConverter) Close(t *Tag) string {
	return "[/" + t.Name + "]"
}

// Tag is a single tage/node in the BBCode tree
type Tag struct {
	Name      string
	Attribute string
	Inner     []*Tag
	Closed    bool
	Parent    *Tag
}

// BBCode returns the tag and its children as a BBCode string
func (t *Tag) BBCode() string {
	return t.Export(BBCodeConverter{})
}

// Export converts the tree to the type specified by the converter
func (t *Tag) Export(c Converter) string {
	if t.Name == "@TEXT@" {
		return t.Attribute
	}
	toRet := ""
	if t.Name != "@BASE@" {
		toRet += c.Open(t)
	}
	for _, tag := range t.Inner {
		toRet += tag.Export(c)
	}
	if t.Name != "@BASE@" {
		toRet += c.Close(t)
	}
	return toRet
}

// Parse parses a BBCode string and generates a tag tree
func Parse(text string) *Tag {
	t := newTokeniser(text)
	baseTag := &Tag{
		Name: "@BASE@",
	}

	currTag := baseTag

	for {
		token, err := t.GetToken()
		if err == io.EOF {
			break
		}
		switch token.Type {
		case tokenText:
			currTag.Inner = append(currTag.Inner, &Tag{
				Name:      "@TEXT@",
				Attribute: token.Data,
			})
		case tokenOpenTag:
			newTag := &Tag{
				Name:   token.Data,
				Parent: currTag,
			}
			currTag.Inner = append(currTag.Inner, newTag)
			currTag = newTag
		case tokenTagAttribute:
			currTag.Attribute = token.Data
		case tokenCloseTag:
			if token.Data != currTag.Name { // Try matching down???
				currTag.Inner = append(currTag.Inner, &Tag{
					Name:      "@TEXT@",
					Attribute: "[/" + token.Data + "]",
				})
			} else {
				currTag.Closed = true
				currTag = currTag.Parent
			}
		case parser.TokenDone:
			break
		}
	}

	return baseTag
}

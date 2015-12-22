package bbcode

import "io"

type Converter interface {
	Open(*Tag) string
	Close(*Tag) string
}

type BBCodeConverter struct{}

func (BBCodeConverter) Open(t *Tag) string {
	if t.Attribute != "" {
		return "[" + t.Name + "=" + t.Attribute + "]"
	}
	return "[" + t.Name + "]"
}

func (BBCodeConverter) Close(t *Tag) string {
	return "[/" + t.Name + "]"
}

type Tag struct {
	Name      string
	Attribute string
	Inner     []*Tag
	Closed    bool
	Parent    *Tag
}

func (t *Tag) BBCode() string {
	return t.Export(BBCodeConverter{})
}

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
		switch token.typ {
		case tokenText:
			currTag.Inner = append(currTag.Inner, &Tag{
				Name:      "@TEXT@",
				Attribute: token.data,
			})
		case tokenOpenTag:
			newTag := &Tag{
				Name:   token.data,
				Parent: currTag,
			}
			currTag.Inner = append(currTag.Inner, newTag)
			currTag = newTag
		case tokenTagAttribute:
			currTag.Attribute = token.data
		case tokenCloseTag:
			if token.data != currTag.Name { // Try matching down???
				currTag.Inner = append(currTag.Inner, &Tag{
					Name:      "@TEXT@",
					Attribute: "[/" + token.data + "]",
				})
			} else {
				currTag.Closed = true
				currTag = currTag.Parent
			}
		case tokenDone:
			break
		}
	}

	return baseTag
}

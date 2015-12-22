package bbcode

import "io"

type Tag struct {
	Name      string
	Attribute string
	Inner     []Tag
	Closed    bool
	Parent    *Tag
}

func (t Tag) BBCode() string {
	return ""
}

func (t Tag) HTML() string {
	return ""
}

func Parse(text string) Tag {
	t := newTokeniser(text)
	baseTag := Tag{
		Name: "@BASE@",
	}

	currTag := &baseTag

	for {
		token, err := t.GetToken()
		if err == io.EOF {
			break
		}
		switch token.typ {
		case tokenText:
			currTag.Inner = append(currTag.Inner, Tag{
				Name:      "@TEXT@",
				Attribute: token.data,
			})
		case tokenOpenTag:
			newTag := Tag{
				Name:   token.data,
				Parent: currTag,
			}
			currTag.Inner = append(currTag.Inner, newTag)
			currTag = &newTag
		case tokenTagAttribute:
			currTag.Attribute = token.data
		case tokenCloseTag:
			if token.data != currTag.Name {
				currTag.Inner = append(currTag.Inner, Tag{
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

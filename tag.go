package bbcode

import (
	"bytes"
	"strings"
	"text/template"
)

// Handler is an interface that represents the text and tag processors.
type Handler interface {
	// Name returns the name of the bbCode tag that this will be used for.
	// Returning an empty string indicates that this Handler should be used
	// for text handling.
	Name() string
	// Handle takes a pointer to the Processor and the attribute to the tag.
	Handle(*Processor, string)
}

// Tag is a simple Handler that just outputs open and closing tags.
type Tag struct {
	name        string
	open, close []byte
}

// NewTag creates a simple Handler that outputs an open and close tag.
// For example, the following would be used to create a tag for handling bold:
// 	NewTag("b", []byte("<b>"), []("</b>"))
func NewTag(name string, open, close []byte) *Tag {
	return &Tag{
		name:  name,
		open:  open,
		close: close,
	}
}

// Name returns the name of the tag.
func (t *Tag) Name() string {
	return t.name
}

// Open outputs the opening of the tag.
func (t *Tag) Open(p *Processor, attr string) {
	p.Write(t.open)
}

// Close outputs the closing of the tag.
func (t *Tag) Close(p *Processor, attr string) {
	p.Write(t.close)
}

// Handle processes the tag.
func (t *Tag) Handle(p *Processor, attr string) {
	t.Open(p, attr)
	p.Process(t.name)
	t.Close(p, attr)
}

// AttrFilterer is used with AttributeTag to provide an attribute filter for
// the AttributeTag. It is used to process the parsed attribute for writing.
type AttrFilterer interface {
	AttrFilter(string) []byte
}

var defaultAttrFilter attrFilter

type attrFilter struct{}

func (attrFilter) AttrFilter(attr string) []byte {
	if attr != "" {
		if !strings.ContainsAny(attr, "'\"&<>\000") {
			return []byte(attr)
		}

		var b bytes.Buffer

		template.HTMLEscape(&b, []byte(attr))

		return b.Bytes()
	}

	return nil
}

// AttrFilterFunc is a wrapper for a func so that it satisfies the AttrFilterer
// interface.
type AttrFilterFunc func(string) []byte

// AttrFilter satisfies the AttrFilterer interface.
func (a AttrFilterFunc) AttrFilter(attr string) []byte {
	return a(attr)
}

// AttributeTag is a simple Handler that outputs and open tag, with an
// attribute, and a close tag.
type AttributeTag struct {
	name                                        string
	open, openClose, attrOpen, attrClose, close []byte
	filter                                      AttrFilterer
}

// NewAttributeTag creates a new Attribute Tag.
// The open byte slice is used to start the open tag and the openClose is used
// to close the open tag.
// The attrOpen and attrClose byte slices are used to surround the filtered
// attribute, within the open tag.
// Lastly, the close byte slice is used for the closing tag.
// The filter is used to modify the attribute, whether to correct formatting
// errors, or to validate. If a nil slice is returned, then no attribute is
// outputted.
// For example the following would create a colour tag for handling font colour:
// 	NewAttributeTag("colour",
// 		[]byte("<span"),
// 		[]byte(">"),
// 		[]byte(" style=\"color: "),
// 		[]byte("\""),
// 		[]byte("</span>"),
//		colourChecker)
//
// A nil filter means that the attr will be written to the output with HTML
// encoding.
func NewAttributeTag(name string, open, openClose, attrOpen, attrClose, close []byte, filter AttrFilterer) *AttributeTag {
	if filter == nil {
		filter = &defaultAttrFilter
	}

	return &AttributeTag{
		name:      name,
		open:      open,
		openClose: openClose,
		attrOpen:  attrOpen,
		attrClose: attrClose,
		close:     close,
		filter:    filter,
	}
}

// Name returns the name of the tag.
func (a *AttributeTag) Name() string {
	return a.name
}

// Open outputs the opening of the tag.
func (a *AttributeTag) Open(p *Processor, attr string) {
	p.Write(a.open)

	if filtered := a.filter.AttrFilter(attr); filtered != nil {
		p.Write(a.attrOpen)
		p.Write(filtered)
		p.Write(a.attrClose)
	}

	p.Write(a.openClose)
}

// Close outputs the closing of the tag.
func (a *AttributeTag) Close(p *Processor, attr string) {
	p.Write(a.close)
}

// Handle processes the tag.
func (a *AttributeTag) Handle(p *Processor, attr string) {
	a.Open(p, attr)
	p.Process(a.name)
	a.Close(p, attr)
}

// OpenClose is an interface for the methods required by FilterTag. Both Tag
// and AttributeTag implement this interface.
type OpenClose interface {
	Name() string
	Open(*Processor, string)
	Close(*Processor, string)
}

// FilterTag is a Handler that filters which child nodes are processed.
type FilterTag struct {
	OpenClose
	filter func(string) bool
}

// NewFilterTag creates a Handler that filters which child nodes are processed.
// The filter takes the name of the tag and returns a bool determining whether
// the tag will be processed as a tag or as text.
// An empty tag name in the filter is used to determine is text is processed.
func NewFilterTag(o OpenClose, filter func(string) bool) *FilterTag {
	return &FilterTag{
		OpenClose: o,
		filter:    filter,
	}
}

// Handle processes the tag, using its filter to determine which children are
// also processed.
func (f *FilterTag) Handle(p *Processor, attr string) {
	f.Open(p, attr)

	allowText := f.filter("")
	name := f.Name()

Loop:
	for {
		switch t := p.Get().(type) {
		case Text:
			if allowText {
				p.Print(t)
			}
		case OpenTag:
			if f.filter(t.Name) {
				p.ProcessTag(t)
			} else if allowText {
				p.Print(t)
			}
		case CloseTag:
			if strings.EqualFold(t.Name, name) {
				break Loop
			}
			if allowText {
				p.Print(p)
			}
		default:
			break Loop
		}
	}

	f.Close(p, attr)
}

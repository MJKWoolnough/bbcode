# bbcode
--
    import "github.com/MJKWoolnough/bbcode"


## Usage

#### type BBCodeConverter

```go
type BBCodeConverter struct{}
```

BBCodeConverter is a Converter that will translate a parsed tree back into
BBCode

#### func (BBCodeConverter) Close

```go
func (BBCodeConverter) Close(t *Tag) string
```
Close creates the closing of a BBCode tag

#### func (BBCodeConverter) Open

```go
func (BBCodeConverter) Open(t *Tag) string
```
Open creates the opening of a BBCode tag

#### type Converter

```go
type Converter interface {
	Open(*Tag) string
	Close(*Tag) string
}
```

Converter is an interface for converting the tag tree into a string

#### type Tag

```go
type Tag struct {
	Name      string
	Attribute string
	Inner     []*Tag
	Closed    bool
	Parent    *Tag
}
```

Tag is a single tage/node in the BBCode tree

#### func  Parse

```go
func Parse(text string) *Tag
```
Parse parses a BBCode string and generates a tag tree

#### func (*Tag) BBCode

```go
func (t *Tag) BBCode() string
```
BBCode returns the tag and its children as a BBCode string

#### func (*Tag) Export

```go
func (t *Tag) Export(c Converter) string
```
Export converts the tree to the type specified by the converter

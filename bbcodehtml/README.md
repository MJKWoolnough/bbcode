# bbcodehtml
--
    import "vimagination.zapto.org/bbcode/bbcodehtml"

This package defines standard bbcode tags.

## Usage

Each of the following is a defined BBCode tag that can be parsed to `bbcode.New`.

|  Export                                                               |  Tag                                                                                                   |
|-----------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| `Align`                                                               | align: can provide an attribute (`left`, `centre`, `center`, `right`, `full`, etc.) to align text.     |
| `LeftAlign`/`CentreAlign`/`CenterAlign`/`RightAlign`/`Full`           | left/centre/center/right/full: used to align text.                                                     |
| `Colour`/`Color`                                                      | colour/color: used to specify a text colour, either as a hex value or a named colour.                  |
| `Font`                                                                | font: used to specify a custom font.                                                                   |
| `Bold`/`Italic`/`Strikethrough`/`Underline`                           | b/i/s/u: used to format text.                                                                          |
| `Size`                                                                | size: allows an attribute between 1 and 50 to set the font size.                                       |
| `Header1`/`Header2`/`Header3`/`Header4`/`Header5`/`Header6`/`Header7` | h1/h2/h3/h4/h5/h6/h7: headers of various depths.                                                       |
| `Quote`                                                               | quote: used to quote text; can provide attr of the quote author.                                       |
| `Code`                                                                | code: used to format code-like text.                                                                   |
| `Image`                                                               | img: used to add images.                                                                               |
| `List`                                                                | list: used to creates lists; inner contents should be enclosed in \[\*\]\[/\*\] tags.                  |
| `Table`                                                               | table: used to create tables; allows the use of \[thead\], \[tr\], \[th\], \[td\], and \[tfoot\] tags. |
| `URL`                                                                 | url: creates a link that points to the attriubte given, wrapping the inner text.                       |

The `All` export combines all of the above into a single slice ready to be used with `bbcode.New`.

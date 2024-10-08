package bbcodehtml

import (
	"strings"
)

// Colours is a map of colour names to their hex equivalent. It is used as an
// attribute filter in an bbcode.AttributeTag where it will recognise the
// contained colours as well as raw hex values.
var Colours = colours{
	"aliceblue":            []byte("#F0F8FF"),
	"antiquewhite":         []byte("#FAEBD7"),
	"aqua":                 []byte("#00FFFF"),
	"aquamarine":           []byte("#7FFFD4"),
	"azure":                []byte("#F0FFFF"),
	"beige":                []byte("#F5F5DC"),
	"bisque":               []byte("#FFE4C4"),
	"black":                []byte("#000000"),
	"blanchedalmond":       []byte("#FFEBCD"),
	"blue":                 []byte("#0000FF"),
	"blueviolet":           []byte("#8A2BE2"),
	"brown":                []byte("#A52A2A"),
	"burlywood":            []byte("#DEB887"),
	"cadetblue":            []byte("#5F9EA0"),
	"chartreuse":           []byte("#7FFF00"),
	"chocolate":            []byte("#D2691E"),
	"coral":                []byte("#FF7F50"),
	"cornflowerblue":       []byte("#6495ED"),
	"cornsilk":             []byte("#FFF8DC"),
	"crimson":              []byte("#DC143C"),
	"cyan":                 []byte("#00FFFF"),
	"darkblue":             []byte("#00008B"),
	"darkcyan":             []byte("#008B8B"),
	"darkgoldenrod":        []byte("#B8860B"),
	"darkgray":             []byte("#A9A9A9"),
	"darkgrey":             []byte("#A9A9A9"),
	"darkgreen":            []byte("#006400"),
	"darkkhaki":            []byte("#BDB76B"),
	"darkmagenta":          []byte("#8B008B"),
	"darkolivegreen":       []byte("#556B2F"),
	"darkorange":           []byte("#FF8C00"),
	"darkorchid":           []byte("#9932CC"),
	"darkred":              []byte("#8B0000"),
	"darksalmon":           []byte("#E9967A"),
	"darkseagreen":         []byte("#8FBC8F"),
	"darkslateblue":        []byte("#483D8B"),
	"darkslategray":        []byte("#2F4F4F"),
	"darkslategrey":        []byte("#2F4F4F"),
	"darkturquoise":        []byte("#00CED1"),
	"darkviolet":           []byte("#9400D3"),
	"deeppink":             []byte("#FF1493"),
	"deepskyblue":          []byte("#00BFFF"),
	"dimgray":              []byte("#696969"),
	"dimgrey":              []byte("#696969"),
	"dodgerblue":           []byte("#1E90FF"),
	"firebrick":            []byte("#B22222"),
	"floralwhite":          []byte("#FFFAF0"),
	"forestgreen":          []byte("#228B22"),
	"fuchsia":              []byte("#FF00FF"),
	"gainsboro":            []byte("#DCDCDC"),
	"ghostwhite":           []byte("#F8F8FF"),
	"gold":                 []byte("#FFD700"),
	"goldenrod":            []byte("#DAA520"),
	"gray":                 []byte("#808080"),
	"grey":                 []byte("#808080"),
	"green":                []byte("#008000"),
	"greenyellow":          []byte("#ADFF2F"),
	"honeydew":             []byte("#F0FFF0"),
	"hotpink":              []byte("#FF69B4"),
	"indianred ":           []byte("#CD5C5C"),
	"indigo ":              []byte("#4B0082"),
	"ivory":                []byte("#FFFFF0"),
	"khaki":                []byte("#F0E68C"),
	"lavender":             []byte("#E6E6FA"),
	"lavenderblush":        []byte("#FFF0F5"),
	"lawngreen":            []byte("#7CFC00"),
	"lemonchiffon":         []byte("#FFFACD"),
	"lightblue":            []byte("#ADD8E6"),
	"lightcoral":           []byte("#F08080"),
	"lightcyan":            []byte("#E0FFFF"),
	"lightgoldenrodyellow": []byte("#FAFAD2"),
	"lightgray":            []byte("#D3D3D3"),
	"lightgrey":            []byte("#D3D3D3"),
	"lightgreen":           []byte("#90EE90"),
	"lightpink":            []byte("#FFB6C1"),
	"lightsalmon":          []byte("#FFA07A"),
	"lightseagreen":        []byte("#20B2AA"),
	"lightskyblue":         []byte("#87CEFA"),
	"lightslategray":       []byte("#778899"),
	"lightslategrey":       []byte("#778899"),
	"lightsteelblue":       []byte("#B0C4DE"),
	"lightyellow":          []byte("#FFFFE0"),
	"lime":                 []byte("#00FF00"),
	"limegreen":            []byte("#32CD32"),
	"linen":                []byte("#FAF0E6"),
	"magenta":              []byte("#FF00FF"),
	"maroon":               []byte("#800000"),
	"mediumaquamarine":     []byte("#66CDAA"),
	"mediumblue":           []byte("#0000CD"),
	"mediumorchid":         []byte("#BA55D3"),
	"mediumpurple":         []byte("#9370DB"),
	"mediumseagreen":       []byte("#3CB371"),
	"mediumslateblue":      []byte("#7B68EE"),
	"mediumspringgreen":    []byte("#00FA9A"),
	"mediumturquoise":      []byte("#48D1CC"),
	"mediumvioletred":      []byte("#C71585"),
	"midnightblue":         []byte("#191970"),
	"mintcream":            []byte("#F5FFFA"),
	"mistyrose":            []byte("#FFE4E1"),
	"moccasin":             []byte("#FFE4B5"),
	"navajowhite":          []byte("#FFDEAD"),
	"navy":                 []byte("#000080"),
	"oldlace":              []byte("#FDF5E6"),
	"olive":                []byte("#808000"),
	"olivedrab":            []byte("#6B8E23"),
	"orange":               []byte("#FFA500"),
	"orangered":            []byte("#FF4500"),
	"orchid":               []byte("#DA70D6"),
	"palegoldenrod":        []byte("#EEE8AA"),
	"palegreen":            []byte("#98FB98"),
	"paleturquoise":        []byte("#AFEEEE"),
	"palevioletred":        []byte("#DB7093"),
	"papayawhip":           []byte("#FFEFD5"),
	"peachpuff":            []byte("#FFDAB9"),
	"peru":                 []byte("#CD853F"),
	"pink":                 []byte("#FFC0CB"),
	"plum":                 []byte("#DDA0DD"),
	"powderblue":           []byte("#B0E0E6"),
	"purple":               []byte("#800080"),
	"rebeccapurple":        []byte("#663399"),
	"red":                  []byte("#FF0000"),
	"rosybrown":            []byte("#BC8F8F"),
	"royalblue":            []byte("#4169E1"),
	"saddlebrown":          []byte("#8B4513"),
	"salmon":               []byte("#FA8072"),
	"sandybrown":           []byte("#F4A460"),
	"seagreen":             []byte("#2E8B57"),
	"seashell":             []byte("#FFF5EE"),
	"sienna":               []byte("#A0522D"),
	"silver":               []byte("#C0C0C0"),
	"skyblue":              []byte("#87CEEB"),
	"slateblue":            []byte("#6A5ACD"),
	"slategray":            []byte("#708090"),
	"slategrey":            []byte("#708090"),
	"snow":                 []byte("#FFFAFA"),
	"springgreen":          []byte("#00FF7F"),
	"steelblue":            []byte("#4682B4"),
	"tan":                  []byte("#D2B48C"),
	"teal":                 []byte("#008080"),
	"thistle":              []byte("#D8BFD8"),
	"tomato":               []byte("#FF6347"),
	"turquoise":            []byte("#40E0D0"),
	"violet":               []byte("#EE82EE"),
	"wheat":                []byte("#F5DEB3"),
	"white":                []byte("#FFFFFF"),
	"whitesmoke":           []byte("#F5F5F5"),
	"yellow":               []byte("#FFFF00"),
	"yellowgreen":          []byte("#9ACD32"),
}

type colours map[string][]byte

func hexFunc(r rune) bool {
	return (r < '0' || r > '9') && (r < 'A' || r > 'F') && (r < 'a' || r > 'f')
}

func (c colours) AttrFilter(attr string) []byte {
	if len(attr) > 0 {
		hex, ok := c[strings.ToLower(attr)]

		if !ok {
			if attr[0] == '#' {
				attr = attr[1:]
			}

			switch len(attr) {
			case 3, 6:
				if strings.IndexFunc(attr, hexFunc) < 0 {
					hex = make([]byte, 1, 7)
					hex[0] = '#'
					hex = append(hex, attr...)
				}
			}
		}

		return hex
	}

	return nil
}

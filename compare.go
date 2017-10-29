package bbcode

import (
	"unicode"
	"unicode/utf8"
)

// Compare preforms a case-insensitive string comparison
func Compare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for {
		ra, size := utf8.DecodeRuneInString(a)
		rb, _ := utf8.DecodeRuneInString(b)
		if ra != rb {
			if unicode.ToLower(ra) != unicode.ToLower(rb) {
				return false
			}
		}
		a = a[size:]
		if len(a) == 0 {
			return true
		}
		b = b[size:]
	}
}

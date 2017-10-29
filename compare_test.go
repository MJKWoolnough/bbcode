package bbcode

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	for n, test := range [...][2]string{
		{"", ""},
		{"", "a"},
		{"a", "a"},
		{"a", "A"},
		{"a", "b"},
		{"a", "aa"},
		{"aa", "aa"},
		{"aA", "Aa"},
		{"è", "è"},
		{"è", "aa"},
		{"Èè", "Èè"},
		{"ÈÈ", "Èè"},
		{"ÈÈ", "èè"},
	} {
		if result := Compare(test[0], test[1]); result != (strings.ToLower(test[0]) == strings.ToLower(test[1])) {
			t.Errorf("test %d: expecting %b, got %b", n+1, !result, result)
		}
	}
}

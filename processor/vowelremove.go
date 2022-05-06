package processor

import "strings"

type vowelRemove struct{}

func newVowelRemove() *vowelRemove {
	return &vowelRemove{}
}

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

// removes vowels from a string
func (v *vowelRemove) manipulate(s string) (result string) {
	for _, c := range s {
		cl := strings.ToLower(string(c))
		if vowels[cl] {
			continue
		} else {
			result = result + string(c)
		}
	}
	return
}

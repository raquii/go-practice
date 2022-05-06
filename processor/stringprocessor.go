package processor

import (
	"strings"
)

type stringProcessor interface {
	manipulate(string) string
}

type reverser struct{}

func newReverser() *reverser {
	return &reverser{}
}

// reverses a string
func (r *reverser) manipulate(s string) (result string) {
	for _, c := range s {
		result = string(c) + result
	}
	return
}

type echoer struct {
	num int
}

func newEchoer() *echoer {
	return &echoer{
		num: 3,
	}
}

// echoes a string
func (e *echoer) manipulate(s string) (result string) {
	x := e.num
	result = s
	for i := 1; i < x; i++ {
		result = result + " " + s
	}
	return
}

type vowelRemover struct{}

func newVowelRemover() *vowelRemover {
	return &vowelRemover{}
}

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

// removes vowels from a string
func (v *vowelRemover) manipulate(s string) (result string) {
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

type mirrorer struct {
	forward bool
}

func newMirrorer() *mirrorer {
	return &mirrorer{
		forward: true,
	}
}

// mirrors a string
func (m *mirrorer) manipulate(s string) (result string) {
	if m.forward {
		for _, c := range s {
			result = string(c) + result
		}
		result = s + result
	} else {
		result = s
		for _, c := range s {
			result = string(c) + result
		}
	}
	return
}

type MyString struct {
	S            string
	Manipulators map[string]stringProcessor
}

func NewMyString(s string) *MyString {
	e := newEchoer()
	r := newReverser()
	v := newVowelRemover()
	m := newMirrorer()
	var processors = map[string]stringProcessor{
		"echo":        e,
		"reverse":     r,
		"vowelremove": v,
		"mirror":      m,
	}
	return &MyString{
		S:            s,
		Manipulators: processors,
	}
}

func (ms *MyString) useManipulation(p stringProcessor) string {
	return p.manipulate(ms.S)
}

func (ms *MyString) UseManipulation(s string) string {
	p := ms.Manipulators[s]
	return ms.useManipulation(p)
}

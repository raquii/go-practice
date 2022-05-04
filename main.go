package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Reverser interface {
	Reverse() string
}

type Echoer interface {
	Echo() string
}

type VowelRemover interface {
	VowelRemove() string
}

type Mirrorer interface {
	Mirror() string
}

type MyString string

// reverses a MyString
func (s MyString) Reverse() (result string) {
	for _, c := range s {
		result = string(c) + result
	}
	return
}

// echoes a MyString
func (s MyString) Echo() string { return string(s) }

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

// removes vowels from a MyString
func (s MyString) VowelRemove() (result string) {
	for _, c := range s {
		if vowels[string(c)] {
			continue
		} else {
			result = string(c) + result
		}
	}
	return
}

// mirrors a MyString
func (s MyString) Mirror() (result string) {
	result = string(s)
	for _, c := range s {
		result = string(c) + result
	}
	return
}

type StringManipulator struct {
	s            MyString
	manipulation map[string]interface{}
}

func (t *StringManipulator) manipulate() {
	t.manipulation = make(map[string]interface{})
	t.manipulation["echo"] = t.s.Echo()
	t.manipulation["reverse"] = t.s.Reverse()
	t.manipulation["mirror"] = t.s.Mirror()
	t.manipulation["vowelremove"] = t.s.VowelRemove()
}

// handles API request routing
func handleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// select the path
		p := r.URL.Path[1:]
		i := strings.Index(p, "/")
		sm := new(StringManipulator)

		// if there is a '/' in `p`, we will see if it is followed by a key in sm.manipulate
		if i > 0 {
			sm.s = MyString(p[:i])
			sm.manipulate()
			fn := strings.ToLower(p[i+1:])
			_, exists := sm.manipulation[fn]
			// if yes, we will render it
			if exists {
				fmt.Fprintf(w, "<h1>%s: %s</h1>", fn, sm.manipulation[fn])
			} else {
				// otherwise we will render the word with a suggestion
				fmt.Fprintf(w, "<h1>%s</h1><h3>pst, try replacing '%s' with one of the manipulators:</h3><ul><li>echo</li><li>mirror</li><li>reverse</li><li>vowelremove</li></ul>", sm.s, fn)
			}
			// if there is no "/" in `p`, we show all of the manipulators
		} else {
			sm.s = MyString(p)
			sm.manipulate()
			for key, val := range sm.manipulation {
				fmt.Fprintf(w, "<h1>%s: %s</h1>", key, val)
			}
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// entrypoint to application
func main() {
	handleRequests()
}

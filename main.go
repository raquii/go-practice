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

type StringManipulator interface {
	Reverser
	Echoer
	VowelRemover
	Mirrorer
}

type MyString string

func (s MyString) Reverse() (result string) {
	for _, c := range s {
		result = string(c) + result
	}
	return
}

func (s MyString) Echo() string { return string(s) }

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

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

func (s MyString) Mirror() (result string) {
	result = string(s)
	for _, c := range s {
		result = string(c) + result
	}
	return
}

// creates an array of manipulated strings
func manipulate(s StringManipulator) [4]string {
	return [4]string{s.Echo(), s.Mirror(), s.Reverse(), s.VowelRemove()}
}

// handles API request routing
func handleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// select the path
		p := r.URL.Path[1:]
		i := strings.Index(p, "/")

		// if there is a '/' in `p`, we will see if it is followed by the name of one of the manipulator functions
		if i > 0 {
			word := MyString(p[:i])
			fn := strings.ToLower(p[i+1:])
			l := manipulate(word)

			funcs := map[string]int{
				"echo":        0,
				"mirror":      1,
				"reverse":     2,
				"vowelremove": 3,
			}

			idx, exists := funcs[fn]
			// if it is one of the manipulator functions, we will render it
			if exists {
				fmt.Fprintf(w, "<h1>%s:</h1><h1>%s</h1>", fn, l[idx])
			} else {
				// otherwise we will render the word with a suggestion
				fmt.Fprintf(w, "<h1>%s</h1><h3>pst, try replacing '%s' with one of the manipulators:</h3><ul><li>echo</li><li>mirror</li><li>reverse</li><li>vowelremove</li></ul>", word, fn)
			}

		} else {
			word := MyString(p)
			l := manipulate(word)
			fmt.Fprintf(w, "<h1>Echo: %s</h1><h1>Mirror: %s</h1><h1>Reverse: %s</h1><h1>VowelRemove: %s</h1>", l[0], l[1], l[2], l[3])
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// entrypoint to application
func main() {
	handleRequests()
}

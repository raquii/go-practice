package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
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

type StringFactory func(str string) interface{}

func createStringHandler(getString StringFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
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

func manipulate(s StringManipulator) {
	fmt.Printf("Echo: %s\n", s.Echo())
	fmt.Printf("Mirror: %s\n", s.Mirror())
	fmt.Printf("Reverse: %s\n", s.Reverse())
	fmt.Printf("VowelRemove: %s\n", s.VowelRemove())
}

// path: /echo/:word
// renders the word input in the URL path in an H1 element
func echoHandler(w http.ResponseWriter, r *http.Request) {
	e := r.URL.Path[len("/echo/"):]
	fmt.Fprintf(w, "<h1>%s</h1>", e)
}

// handles API request routing
func handleRequests() {
	http.HandleFunc("/echo/", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// entrypoint to application
func main() {
	// handleRequests()
	m := MyString("This is a test")
	re := regexp.MustCompile(`([^/]+)/?$`)
	r := "/test/what/iwantthis"
	fmt.Printf("%q\n", re.FindString(r))
	manipulate(m)
}

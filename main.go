package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	handleRequests()
}

package main

import (
	"example/getgoing/duck"
	"example/getgoing/processor"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// handles API request routing
func handleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// select the path
		p := r.URL.Path[1:]
		i := strings.Index(p, "/")

		// if there is a '/' in `p`, we will see if it is followed by a key in sm.manipulate
		if i > 0 {
			sm := processor.NewMyString(p[:i])
			fn := strings.ToLower(p[i+1:])
			_, exists := sm.Manipulators[fn]
			// if yes, we will render it
			if exists {
				fmt.Fprintf(w, "<h1>%s: %s</h1>", fn, sm.UseManipulation(fn))
			} else {
				// otherwise we will render the word with a suggestion
				fmt.Fprintf(w, "<h1>%s</h1><h3>pst, try replacing '%s' with one of the manipulators:</h3><ul><li>echo</li><li>mirror</li><li>reverse</li><li>vowelremove</li></ul>", sm.S, fn)
			}
			// if there is no "/" in `p`, we show all of the manipulators
		} else {
			sm := processor.NewMyString(p)
			for key := range sm.Manipulators {
				fmt.Fprintf(w, "<h1>%s: %s</h1>", key, sm.UseManipulation(key))
			}
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// entrypoint to application
func main() {
	duck.DuckTime()
}

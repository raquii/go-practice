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
		sp := processor.NewProcessor()

		// if there is a '/' in `p`, we will see if it is followed by a key in `m`
		if i > 0 {
			word := p[:i]

			sm := sp.UseProcessors(word)
			fn := strings.ToLower(p[i+1:])
			_, exists := sm[fn]
			// if yes, we will render it
			if exists {
				fmt.Fprintf(w, "<h1>%s: %s</h1>", fn, sm[fn])
			} else {
				// otherwise we will render the word with a suggestion
				fmt.Fprintf(w, "<h1>%s</h1><h3>pst, try replacing '%s' with one of the manipulators:</h3><ul><li>echo</li><li>mirror</li><li>reverse</li><li>vowelremove</li></ul>", word, fn)
			}
			// if there is no "/" in `p`, we show all of the manipulators
		} else {
			sm := sp.UseProcessors(p)
			for manipulator, manipulation := range sm {
				fmt.Fprintf(w, "<h1>%s: %s</h1>", manipulator, manipulation)
			}
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func yesOrNo() (b bool) {
	var i string
	fmt.Println(" y/n")

	for {
		fmt.Scan(&i)

		if i == "n" {
			b = false
			break
		} else if i == "y" {
			b = true
			break
		}
		fmt.Println("Invalid!: Enter y for 'yes', n for 'no'.")
	}
	return
}

func startDucks() {
	fmt.Print("Would you like to make a duck?")
	r := yesOrNo()

	if r {
		var n string
		fmt.Println("\nWhat kind of duck will you make?")
		fmt.Scan(&n)
		fmt.Printf("\nNice, a %s duck\n", n)

		fmt.Printf("\nCan %s duck fly?", n)
		f := yesOrNo()

		fmt.Printf("\nCan %s duck quack?", n)
		q := yesOrNo()

		d := duck.NewDuck(n, f, q)

		d.ActLikeDuck()
		startDucks()
	} else {
		fmt.Print("Okay. Quitting Ducks...")
		startServer()
	}
}

func startServer() {
	fmt.Print("Would you like to start the server?")
	r := yesOrNo()

	if r {
		fmt.Print("Starting...visit: localhost:8080/hello")
		handleRequests()
	} else {
		fmt.Print("Goodbye")
	}
}

// entrypoint to application
func main() {
	startDucks()
}

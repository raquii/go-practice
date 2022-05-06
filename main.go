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

func start() {
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
		start()
	} else {
		fmt.Print("Okay. Quitting...")
	}
}

// entrypoint to application
func main() {
	start()
}

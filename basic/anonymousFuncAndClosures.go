package basic

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sort"
)

func AnonymousFunctionsExample() {
	// func literal to create anonymous function
	// func(){...}()
	func() {
		fmt.Println("hello from anonymous function!")
	} ()
	
	// store function in a var
	// func literal not executed
	myFunc := func() int {
		fmt.Println("I'll return number 5")
		return 5
	}
	fmt.Println(reflect.TypeOf(myFunc))

	// func literal invoked
	myFuncInv := func() int {
		fmt.Println("I'll return number 10")
		return 10
	} ()
	fmt.Println(reflect.TypeOf(myFuncInv))

	// function types
	// function type designates all functions with the same parameters and results
	type Funky func(string)
	var f Funky
	f = func(s string) {
		log.Printf("Funky %s", s)
	}
	f("bit")
}

func ClosuresExample() {
	// printer assigned to variable
	p := printer()
	p()
	p()
	p()

	// function params example
	scores := []int{99, 10, 54, 77, 15, 84, 11}

	// ascending
	sort.Slice(scores, func(i, j int) bool {return scores[i] < scores[j]})
	fmt.Println(scores)
	// descending
	sort.Slice(scores, func(i, j int) bool {return scores[i] > scores[j]})
	fmt.Println(scores)
}

// closure example
func printer() func() {
	k := 1
	return func() {
		fmt.Printf("Print #%v\n", k)
		k++
	}
}

func BasicWebServer() {
	http.HandleFunc("/homepage", homepageHandler)
	// handler with warapper technique
	http.HandleFunc("/wrapped", trackVisits(wrappedHandler))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage\n")
	fmt.Fprintf(w, "Nice to see you here\n")
}

func wrappedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm wrapped by visit tracker\n")
}

func trackVisits(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// track the visit
		fmt.Println("one visit!")
		// call original handler
		handler(w, r)
	}
}

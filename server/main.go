package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is the about page.")
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	fmt.Fprintf(w, "Hello, %s!", name)
// }

func main() {
	http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/about", aboutHandler)
	// http.HandleFunc("/hello", handler) // custom handler for "/hello" route

	//http.HandleFunc("/", userHandler)
	fmt.Print("Openning server on http://localhost:8000\n")
	http.ListenAndServe(":8000", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/huh" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "brother this isn't supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "wassup bro")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	goon := r.FormValue("goon")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Goon = %s\n", goon)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //tell go to check static folder. looks at index.html first cuz that's just how it be.
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/huh", helloHandler)

	fmt.Println("Starting server at port 8080, ok?")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

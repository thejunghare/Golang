package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/server" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Server started")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprint(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	/* Response form server */
	http.HandleFunc("/server", serverHandler)

	fmt.Println("Starting server at port 8080")

	/* start a web server */
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

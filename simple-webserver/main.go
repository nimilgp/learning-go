package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprintf(w, "NAME:%s\n", name)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}

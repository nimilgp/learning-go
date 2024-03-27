package main

import (
    "net/http"
	"fmt"
)

func getHello(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
    }

	func getEcho(w http.ResponseWriter, r *http.Request) {
		val := r.PathValue("val")
		fmt.Fprintf(w, "%s\n", val)
    }

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/",getHello)
	mux.HandleFunc("GET /echo/{val}",getEcho)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Println("port not availible or some other error")
	}
}

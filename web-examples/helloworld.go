package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>hello world!!</h1>")
    })

    http.ListenAndServe(":3333", nil)
}

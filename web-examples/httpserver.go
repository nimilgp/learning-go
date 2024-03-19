package main

import (
    "fmt"
    "net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>ROOT PAGE</h1> welcome to my website!!!<br><a href='/leaf1'>LEAF1</a> <a href='/leaf2'>LEAF2</a>")
    }

func getLeaf1(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>THIS IS LEAF NODE 1</h1><a href='/'>root</a>")
    }

func getLeaf2(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>THIS IS LEAF NODE 2</h1><a href='/'>root</a>")
    }

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/leaf1", getLeaf1)
    http.HandleFunc("/leaf2", getLeaf2)

    http.ListenAndServe(":3333", nil)
}

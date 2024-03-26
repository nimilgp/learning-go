package main

import (
	"net/http"
	"fmt"
)

func echoVal(w http.ResponseWriter, r *http.Request) {
	val := r.PathValue("val")
	w.Write([]byte(val + "\n"))
}

func echoWhy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("to echo back\n"))
}

func addUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	number := r.PathValue("number")
	fmt.Println("username:",username,"\nnumber:",number)
}


func main() {
	router := http.NewServeMux()

	router.HandleFunc("/echo/{val}", echoVal)
	router.HandleFunc("/echo/why", echoWhy)

	//curl -X POST http://localhost:3333/adduser/nimil/65
	//only POST REQUEST is allowed
	router.HandleFunc("POST /adduser/{username}/{number}", addUser)

	http.ListenAndServe(":3333", router)
}

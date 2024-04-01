package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func gotName(w http.ResponseWriter, r *http.Request) {
	nameVal := r.PathValue("nameVal")
	w.Write([]byte("name was: " + nameVal))
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /name/{nameVal}", gotName)

	v2 := http.NewServeMux()
	v2.Handle("/v2/", http.StripPrefix("/v2", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Starting Server at : %s", s.addr)

	return server.ListenAndServe()
}

package main

import "log"

func main() {
	server := NewApiServer(":3333")

	err := server.Run()
	if err != nil {
		log.Fatal(err, " in the ListAndServe call")
	}
}

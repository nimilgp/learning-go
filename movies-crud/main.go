package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	movies = append(movies, Movie{ID: "1", Isbn: "2300", Title: "aot", Director: &Director{Firstname: "hajime", Lastname: "iseyama"}})
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

var movies []Movie

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "2300", Title: "aot", Director: &Director{Firstname: "hajime", Lastname: "iseyama"}})
	movies = append(movies, Movie{ID: "4", Isbn: "4900", Title: "op", Director: &Director{Firstname: "eichiro", Lastname: "oda"}})
	movies = append(movies, Movie{ID: "6", Isbn: "8230", Title: "zelda", Director: &Director{Firstname: "shigeru", Lastname: "miyamoto"}})
	movies = append(movies, Movie{ID: "7", Isbn: "9040", Title: "hk", Director: &Director{Firstname: "ari", Lastname: "william"}})
	movies = append(movies, Movie{ID: "8", Isbn: "9300", Title: "sotc", Director: &Director{Firstname: "fumeto", Lastname: "ueda"}})
	router := http.NewServeMux()

	router.HandleFunc("GET /movies", getMovies)
	router.HandleFunc("GET /movies/{id}", getMovie)
	router.HandleFunc("POST /movies", createMovie)
	// router.HandleFunc("PUT /movies/{id}",updateMovies)
	router.HandleFunc("DELETE /movies/{id}", deleteMovies)

	fmt.Println("server at port 3333")
	log.Fatal(http.ListenAndServe(":3333", router))
}

// eg: curl -X DELETE http://localhost:3333/movies/4

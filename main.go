package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // director is of type director
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// creating a variable of movies of type slice Movie
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // content type to json
	json.NewEncoder(w).Encode(movies)                  // encode movies to json and send to frontend
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// after deleting movie return the all remaining movies
	json.NewEncoder(w).Encode(movies)
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// decode the body
	_ = json.NewDecoder(r.Body).Decode(&movie) // we will get the body after decoding in var movie declare just above
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop over movie range
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")
	params := mux.Vars(r) // params has the id
	// inside the movies finding the id and if match it then delete it.
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	r := mux.NewRouter() // r is the new router

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "Movie is Movie",
		Director: &Director{
			Firstname: "James",
			Lastname:  "Cameron",
		}})
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "5678",
		Title: "Not a Movie",
		Director: &Director{
			Firstname: "Mashiyat",
			Lastname:  "Hussain",
		},
	})

	// routes with their function and their methods
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	// starting the server
	fmt.Printf("Starting the server at 8000...\n")

	log.Fatal(http.ListenAndServe(":8000", r))
}

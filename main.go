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

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Poster   string    `json:"poster"` // URL de la imagen de la película
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Image     string `json:"image"` // URL de la imagen del director
}

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
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
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, movie{
		ID:     "1",
		Isbn:   "31311",
		Title:  "Titanic",
		Poster: "https://static.nationalgeographicla.com/files/styles/image_3200/public/nationalgeographic762774.jpg?w=1600&h=900", // URL de la imagen de la película
		Director: &Director{
			Firstname: "Alan",
			Lastname:  "Doe",
			Image:     "https://upload.wikimedia.org/wikipedia/commons/7/77/James_Cameron_October_2012.jpg", // URL de la imagen del director
		},
	})

	movies = append(movies, movie{
		ID:     "2",
		Isbn:   "34111",
		Title:  "Harry Potter",
		Poster: "https://www.futuro.cl/wp-content/uploads/2024/02/Harry-Potter-768x432.webp", // URL de la imagen de la película
		Director: &Director{
			Firstname: "Juan",
			Lastname:  "Carrera",
			Image:     "https://s3.ppllstatics.com/elcorreo/www/multimedia/2023/09/28/actor-harry-potter-k2GF-U210277961782Yd-1200x840@El%20Correo.PNG", // URL de la imagen del director
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

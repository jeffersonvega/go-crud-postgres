package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	l "go-crud/DB"
	"log"
	"net/http"
)

type Movie struct {
	MovieID   string `json:"id"`
	MovieName string `json:"title"`
	MovieYear string `json:"year"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/movies/", GetMovies).Methods("GET")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("accediendGetMovies")
	db := l.SetupDB()

	printMessage("Getting movies...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM movies")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var movies []Movie

	// Foreach movie
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		checkErr(err)

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = JsonResponse{Type: "success", Data: movies}

	json.NewEncoder(w).Encode(response)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

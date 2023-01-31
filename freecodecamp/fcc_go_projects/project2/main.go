package main

import "github.com/gorilla/mux"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbin"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	//go get "github.com/gorilla/mux" to download this package
	// method is now deprecated use `go install <package>`
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies)

}

func getMovies() {

}

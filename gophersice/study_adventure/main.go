package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smiggiddy/golang/gophercise/study_adventure/studyguide"
)

type Page struct {
	Title string
	Body  []byte
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there!", r.URL.Path)
}

// func storyHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf("testing")
// }

func main() {
	studyData := studyguide.Parse()

	fmt.Println(studyData)
	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/story", storyHandler)
	log.Fatal(http.ListenAndServe(":5080", nil))

}

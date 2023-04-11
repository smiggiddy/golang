package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smiggiddy/golang/gophercise/study_adventure/studyguide"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// sg := studyguide.Parse("./studyguide.json")
	// t, _ := template.ParseFiles("templates/main.html")
	// t.Execute(w, sg)

	// studyguide.RenderTemplate(w, r, "main", sg)
	fmt.Fprintf(w, "Welcome to the stuff")

}

// func storyHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf("testing")
// }

func main() {
	// studyData := studyguide.Parse()

	// var studyTopics studyguide.Topics

	studyTopics := studyguide.Parse("./studyguide.json")

	fmt.Println(studyTopics)
	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/story", storyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

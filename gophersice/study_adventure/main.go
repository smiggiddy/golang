package main

import (
	"log"
	"net/http"

	"github.com/smiggiddy/golang/gophercise/study_adventure/studyguide"
)

type Page struct {
	Title string
	Body  struct{}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	sg := studyguide.Parse()
	// t, _ := template.ParseFiles("templates/main.html")
	// t.Execute(w, sg)

	studyguide.RenderTemplate(w, r, "main", sg)

}

// func storyHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf("testing")
// }

func main() {
	// studyData := studyguide.Parse()

	// fmt.Println(studyData)
	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/story", storyHandler)
	log.Fatal(http.ListenAndServe(":5080", nil))

}

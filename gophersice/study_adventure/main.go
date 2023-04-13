package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/smiggiddy/golang/gophercise/study_adventure/studyguide"
)

func renderTemplate(w http.ResponseWriter, tmpl string, s studyguide.Study) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, s)
}

func homeHandler(w http.ResponseWriter, r *http.Request, t studyguide.Topics) {
	renderTemplate(w, "main", t["start"])
}

func cksHandler(w http.ResponseWriter, r *http.Request, t studyguide.Topics) {
	renderTemplate(w, "main", t["cks"])
}

func cksvimHandler(w http.ResponseWriter, r *http.Request, t studyguide.Topics) {
	renderTemplate(w, "main", t["cks-vim"])
}

func ckslinksHandler(w http.ResponseWriter, r *http.Request, t studyguide.Topics) {
	renderTemplate(w, "main", t["cks-links"])
}

func awsHandler(w http.ResponseWriter, r *http.Request, t studyguide.Topics) {
	renderTemplate(w, "main", t["aws"])
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, studyguide.Topics)) http.HandlerFunc {
	sg := studyguide.Parse("./studyguide.json")
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		fn(w, r, sg)
	}

}

func main() {
	http.HandleFunc("/", makeHandler(homeHandler))
	http.HandleFunc("/cks", makeHandler(cksHandler))
	http.HandleFunc("/cks-vim", makeHandler(cksvimHandler))
	http.HandleFunc("/cks-links", makeHandler(ckslinksHandler))
	http.HandleFunc("/aws", makeHandler(awsHandler))
	// http.HandleFunc("/story", storyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

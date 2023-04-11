package studyguide

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Parse(file string) Topics {
	var Data Topics

	// b, err := os.ReadFile(file)
	// if err != nil {
	// 	fmt.Println("Something went wrong with that file boss")
	// }

	// // var f interface{}
	// var F studyData
	// jsonERR := json.Unmarshal(b, &F)
	// if jsonERR != nil {
	// 	fmt.Println("Something went wrong with that file boss")
	// }
	// return F

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file")
	}
	d := json.NewDecoder(f)

	if err := d.Decode(&Data); err != nil {
		fmt.Println(err)
	}
	return Data
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, s studyData) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, s)
}

type Topics map[string]Study

type Study struct {
	Title   string         `json:"title"`
	Desc    string         `json:"desc"`
	Tips    []string       `json:"tips"`
	Options []StudyOptions `json:"options"`
}

type StudyOptions struct {
	Text string `json:"text"`
	Sect string `json:"sect"`
}

type studyData struct {
	Intro struct {
		Title   string   `json:"title"`
		Desc    string   `json:"desc"`
		About   []string `json:"tips"`
		Options []struct {
			Text string `json:"text"`
			Sect string `json:"sect"`
		} `json:"options"`
	} `json:"start"`
	CKS struct {
		Title   string   `json:"title"`
		Desc    string   `json:"desc"`
		About   []string `json:"tips"`
		Options []struct {
			Text string `json:"text"`
			Sect string `json:"sect"`
		} `json:"options"`
	} `json:"cks"`
	AWS struct {
		Title   string   `json:"title"`
		Desc    string   `json:"desc"`
		About   []string `json:"tips"`
		Options []struct {
			Text string `json:"text"`
			Sect string `json:"sect"`
		} `json:"options"`
	} `json:"aws"`
	CKSVim struct {
		Title   string   `json:"title"`
		Desc    string   `json:"desc"`
		About   []string `json:"tips"`
		Options []struct {
			Text string `json:"text"`
			Sect string `json:"sect"`
		} `json:"options"`
	} `json:"cks-vim"`
	CKSLinks struct {
		Title   string   `json:"title"`
		Desc    string   `json:"desc"`
		About   []string `json:"tips"`
		Options []struct {
			Text string `json:"text"`
			Sect string `json:"sect"`
		} `json:"options"`
	} `json:"cks-links"`
}

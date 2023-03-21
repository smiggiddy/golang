package studyguide

import (
	"encoding/json"
	"fmt"
	"os"
)

func Parse() {
	var b []byte

	b, err := os.ReadFile("/Users/smig/repos/github-golang/gophersice/study_adventure/studyguide.json")
	if err != nil {
		fmt.Println("Something went wrong with that file boss")
	}

	// var f interface{}
	var f studyData
	jsonERR := json.Unmarshal(b, &f)
	if jsonERR != nil {
		fmt.Println("Something went wrong with that file boss")
	}
	fmt.Println(f.Intro.Title)

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

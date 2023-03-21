package studyguide

import (
	"encoding/json"
	"fmt"
	"os"
)

func Parse() studyData {
	var b []byte

	b, err := os.ReadFile("./studyguide.json")
	if err != nil {
		fmt.Println("Something went wrong with that file boss")
	}

	// var f interface{}
	var F studyData
	jsonERR := json.Unmarshal(b, &F)
	if jsonERR != nil {
		fmt.Println("Something went wrong with that file boss")
	}
	return F
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

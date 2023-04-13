package studyguide

import (
	"encoding/json"
	"fmt"
	"os"
)

func Parse(file string) Topics {
	var Data Topics

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

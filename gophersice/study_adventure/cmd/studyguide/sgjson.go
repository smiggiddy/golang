package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/smiggiddy/golang/gophercise/study_adventure/studyguide"
)

func main() {

	filename := flag.String("file", "studyguide.json", "JSON file path to parse")

	flag.Parse()

	fmt.Printf("Parsing the JSON file: %s\n", *filename)

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)

	var story studyguide.Topics

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
}

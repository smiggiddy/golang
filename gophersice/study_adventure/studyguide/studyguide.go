package studyguide

import (
	"encoding/json"
	"fmt"
	"os"
)

func Parse() {
	fmt.Println("This is just some starter code")
	var b []byte

	b, err := os.ReadFile("/home/smig/repos/gophercises/study_adventure/yamlvalues.yaml")
	if err != nil {
		fmt.Println("Something went wrong with that file boss")
	}

	var f interface{}

	jsonERR := json.Unmarshal(b, &f)
	if jsonERR != nil {
		fmt.Println("Something went wrong with that file boss")
	}
	fmt.Println(f)
}

package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string //slice of strings
}

func main() {
	fmt.Println("Maps and Stuff in Go")

	learningTopics := map[string]bool{
		"Golang":     true,
		"Python":     true,
		"Javascript": false,
	}

	fmt.Println(learningTopics)

	// declaring a map
	m := make(map[string]int)

	m["ghosts"] = 3

	linuxIs := map[string]string{}

	linuxIs["fun"] = "bash"

	fmt.Println(linuxIs)
	//remove from map

	delete(linuxIs, "fun")

	// test if key exists
	_, ok := linuxIs["fun"]

	fmt.Println(ok)

	// maps assign by reference like a slice
	//changing in one place will change everywhere

	//instantiate a struct, use field names in the instantiate
	aDoc := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"This dude",
			"That girl",
		},
	}
	fmt.Println(aDoc)
	fmt.Println(aDoc.actorName)

	// go exports variables that start with a capital letter to the package
	// lowercase internal to the package
	// use pascal casing and camel casing

	// good for a quick ops, but
	aTest := struct{ name string }{name: "smig"}

	fmt.Println(aTest)
}

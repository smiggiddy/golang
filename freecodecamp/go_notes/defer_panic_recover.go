package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	// executes after main function returns
	defer fmt.Println("middle")

	fmt.Println("end")

	// defer closes out like a reverse queue LIFO

	res, err := http.Get("http://smig.tech/robots.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
}

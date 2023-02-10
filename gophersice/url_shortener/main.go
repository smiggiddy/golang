package main

import (
	"fmt"
	"log"
	"urlshortener/urlshortener"

	"gopkg.in/yaml.v3"
)

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	t := T{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
	fmt.Println("test")
	urlshortener.PrintHello()
	// 	mux := defaultMux()

	// 	// Build the MapHandler using the mux as the fallback
	// 	pathsToUrls := map[string]string{
	// 		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// 	}
	// 	mapHandler := url_shortener.MapHandler(pathsToUrls, mux)

	//	// Build the YAMLHandler using the mapHandler as the
	//	// fallback
	//	yaml := `
	//   - path: /urlshort
	//     url: https://github.com/gophercises/urlshort
	//   - path: /urlshort-final
	//     url: https://github.com/gophercises/urlshort/tree/solution
	//
	// `
	//
	//	yamlHandler, err := url_shortener.YAMLHandler([]byte(yaml), mapHandler)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("Starting the server on :8080")
	//	http.ListenAndServe(":8080", yamlHandler)
}

// func defaultMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	return mux
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world!")
// }

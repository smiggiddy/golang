package main

import (
	"flag"
	"fmt"
	"net/http"
	"urlshortener/urlshortener"
)

func main() {
	var y []byte
	filePtr := flag.String("yaml", "", "parse url values from a yaml file")
	mux := defaultMux()
	flag.Parse()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/smig":           "https://smig.tech",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	yaml := `
      - path: /urlshort
        url: https://github.com/gophercises/urlshort
      - path: /urlshort-final
        url: https://github.com/gophercises/urlshort/tree/solution
    `

	yamlData, err := urlshortener.YamlReader(*filePtr)

	if err != nil {
		y = []byte(yaml)
		fmt.Println("This ran cause of the nil")
	}

	if yamlData != nil {
		y = yamlData
	}

	yamlHandler, err := urlshortener.YAMLHandler(y, mapHandler)

	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

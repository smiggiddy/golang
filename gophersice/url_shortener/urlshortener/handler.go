package urlshortener

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type redirects struct {
	url       string
	shortName string
}

type yamlURLs struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func redirector(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusFound)
}

func buildMap(fn func(w http.ResponseWriter, r *http.Request, u string), url map[string]string, fb http.Handler) http.HandlerFunc {

	re := []redirects{}

	for s, u := range url {
		re = append(re, redirects{
			url:       u,
			shortName: s,
		})
	}
	urlMap := func(w http.ResponseWriter, r *http.Request) {
		for _, u := range re {
			if r.URL.Path == u.shortName {
				fn(w, r, u.url)
			}
		}
		fb.ServeHTTP(w, r)
	}

	return urlMap
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	var bm http.HandlerFunc
	bm = buildMap(redirector, pathsToUrls, fallback)

	return bm
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...

	urls := make([]yamlURLs, 2)
	err := yaml.Unmarshal(yml, &urls)

	if err != nil {
		fmt.Printf("An error occured %v\n", err)
	}

	urlMaps := make(map[string]string)

	for _, v := range urls {
		urlMaps[v.Path] = v.Url
	}

	var bm http.HandlerFunc
	bm = buildMap(redirector, urlMaps, fallback)

	return bm, err
}

func PrintHello() {
	fmt.Println("Testing")
}

func YamlReader(file string) ([]byte, error) {
	var y []byte
	var test yamlURLs
	yamlFile, err := os.Open(file)

	if err != nil {
		fmt.Println("An error occured with the file ::", err)
	}

	reader := yaml.NewDecoder(yamlFile)

	if err := reader.Decode(&test); err != nil {
		return nil, errors.New("Error with yaml")
	}
	fmt.Println("Hello")
	fmt.Println(test)
	return y, nil
}

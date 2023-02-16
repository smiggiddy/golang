package urlshortener

import (
	"errors"
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

type redirects struct {
	url       string
	shortName string
}

func redirectMap(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusFound)
}

func fallBack(w http.ResponseWriter, r *http.Request) {

}

var found = false

func buildMap(fn func(w http.ResponseWriter, r *http.Request, u string), url map[string]string) (http.HandlerFunc, error) {

	re := []redirects{}

	for s, u := range url {
		re = append(re, redirects{
			url:       u,
			shortName: s,
		})
	}
	a := func(w http.ResponseWriter, r *http.Request) {
		for _, u := range re {
			if r.URL.Path == u.shortName {
				fn(w, r, u.url)
				found = true
			}
		}

	}
	if found == true {
		return a, nil
	}
	return nil, errors.New("go to default")
}

func (re redirects) checkPath(r *http.Request, p string) bool {
	return r.URL.Path == p
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	var bm http.HandlerFunc
	var err error

	bm, err = buildMap(redirectMap, pathsToUrls)
	if err != nil {
		fmt.Println("This block ran ")
		return fallBack
	}

	return bm // fallback.ServeHTTP
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}

func PrintHello() {
	fmt.Println("Testing")
}

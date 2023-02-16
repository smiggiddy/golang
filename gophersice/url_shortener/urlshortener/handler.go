package urlshortener

import (
	"fmt"
	"net/http"
)

type redirects struct {
	url       string
	shortName string
}

func redirectMap(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusFound)
}

func fallBack(w http.ResponseWriter, r *http.Request) {

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

func (re redirects) checkPath(r *http.Request, p string) bool {
	return r.URL.Path == p
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	var bm http.HandlerFunc
	bm = buildMap(redirectMap, pathsToUrls, fallback)

	return bm
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

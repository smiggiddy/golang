package urlshortener

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

func printThis(s string) {
	fmt.Printf("%v", s)
}

type redirects struct {
	url       string
	shortName string
}

func (re *redirects) redirectMap(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, re.url, http.StatusFound)
}

func buildMap(fn func(w http.ResponseWriter, r *http.Request), sn string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == sn {
			fn(w, r)
		}
	}

}

func (re redirects) checkPath(r *http.Request, p string) bool {
	return r.URL.Path == p
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	// mux := func(w http.ResponseWriter, r *http.Request) {
	// 	// http.Redirect(w, r, u, 200)
	// 	fmt.Fprintf(w, "a field")
	// }

	for p, u := range pathsToUrls {

		re := redirects{
			url:       u,
			shortName: p,
		}

		bm := buildMap(re.redirectMap, re.shortName)

		// if err != nil {
		// 	continue
		// }
		return bm
		// if re.checkPath(r, p) {
		// 	return buildMap(re.redirectMap, re.shortName)
		// } else {
		// 	continue
		// }
		// http.HandleFunc(re.shortName, re.buildMap)
		// return re.buildMap
		// http.RedirectHandler(u, 302)
	}
	return fallback.ServeHTTP
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

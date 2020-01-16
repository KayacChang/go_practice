package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathMap map[string]string

func main() {
	mux := newMux()

	yamlByte := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	handler, err := YAMLHandler([]byte(yamlByte), mux)

	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", handler)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func MapHandler(pathsToUrls pathMap, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		// if we can match a path
		if dest, ok := pathsToUrls[path]; ok {
			// redirect to it
			http.Redirect(w, r, dest, http.StatusFound)

			return
		}

		// else fallback
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlByte []byte, fallback http.Handler) (http.HandlerFunc, error) {

	type pathURL struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}

	// 1. Parse the yaml
	var pathUrls []pathURL

	err := yaml.Unmarshal(yamlByte, &pathUrls)

	if err != nil {
		return nil, err
	}

	// 2. Convert YAML array into map
	paths := pathMap{}

	for _, pu := range pathUrls {
		paths[pu.Path] = pu.URL
	}

	// 3. return map handler
	return MapHandler(paths, fallback), nil
}

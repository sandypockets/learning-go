package main

import (
	"fmt"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if the path is in the map, then redirect to the url
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func main() {
	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/blog":       "https://sandypockets.dev",
		"/github":     "https://github.com",
		"/twitter":    "https://twitter.com",
	}

	mapHandler := MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", mapHandler)
	if err != nil {
		return 
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/bye", goodbye)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye!")
}

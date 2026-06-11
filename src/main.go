package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
)

var store = make(map[string]string)

func main() {
	port := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	// takes a regular URL to shorten
	mux.HandleFunc("POST /shorten", handleShorten)
	// takes a shortened URL to redirect
	mux.HandleFunc("GET /{slug}", handleRedirect)

	fmt.Printf("[INFO] server started on port %s\n", port)
	http.ListenAndServe(port, mux)
}

func handleRoot(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintf(w, "ok")
}

func handleShorten(
	w http.ResponseWriter,
	r *http.Request,
) {
	// store the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	// convert byte slice body to a string
	url := string(body)

	// convert URL to zwc string
	urlShortened := encode(hash(url))

	fmt.Fprintf(w, urlShortened)

	store[urlShortened] = url
	fmt.Printf("[INFO] stored %s\n", url)
}

func handleRedirect(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Printf("[INFO] %s %s\n", r.Method, r.URL.Path)
	// extract the value at path slug into a variable
	slug := r.PathValue("slug")

	// find slug in the store and return url and a boolean ok for status
	url, ok := store[slug]
	fmt.Printf("[INFO] fetching url for slug in store\n")

	if !ok {
	    fmt.Printf("[ERROR] slug not found in store\n")
	    http.NotFound(w, r)
	    return
	}
	fmt.Printf("[INFO] slug found, redirecting to %s\n", url)
	// redirect with a 301 to cache the redirect permenantly
	// because urls are hashed, the same url always returns the same hash
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func hash(s string) uint32 {
	hash := fnv.New32()
	hash.Write([]byte(s))
	return hash.Sum32()
}

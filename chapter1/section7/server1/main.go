// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Counter : Global counter variable
var Counter int

func main() {
	http.HandleFunc("/count", countHandler) // each request calls handler
	http.HandleFunc("/", handler)           // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Scheme = %s\n", r.URL.Scheme)
	fmt.Fprintf(w, "URL.Host = %s\n", r.URL.Host)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "RemoteAddr = %s\n", r.RemoteAddr)
}

// handler echoes the Path component of the requested URL.
func countHandler(w http.ResponseWriter, r *http.Request) {
	Counter++
	fmt.Fprintf(w, "count: %d\n", Counter)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

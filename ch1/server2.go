// server2 is a minimal echo and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request){
	mu.Lock() // ensure at most one goroutine access the counter
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

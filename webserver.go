package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// The r.URL.Path[1:] part tells the string to read the first 
	// 		part of the URL after the main domain address
	// This makes the webserver quite dynamic
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// Initiate the webserver
    http.HandleFunc("/", handler)
    // Run and listen on port 8080 until the service is stopped
    http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request received!")
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	log.Print("Server running on port 8080!")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello, World!</html></body></h1>")
}

func main() {
	// serves all url inlcude /up required by kamal
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

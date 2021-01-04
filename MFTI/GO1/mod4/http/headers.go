package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// set some Header
	w.Header().Set("RequestID", "lyalyalya...lya")

	// get UserAgent info
	fmt.Fprintf(w, "You browser is: %s\n\n", r.UserAgent())
	fmt.Fprintf(w, "You accept %s\n", r.Header.Get("Accept"))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

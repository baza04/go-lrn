package main

import (
	"fmt"
	"net/http"
)

// query params can be parsed by few ways
func handler(w http.ResponseWriter, r *http.Request) {
	// 1st way from URL by Query func (only for GET method)
	myParam := r.URL.Query().Get("param")
	if myParam != "" { // return "" if not found parameter
		fmt.Fprintf(w, "`myParam` is %s\n", myParam)
	}

	// 2nd way from r (Request object) by FromValue func
	key := r.FormValue("key") // also can be user to ResponseWriter
	if key != "" {            // return "" if not found parameter
		fmt.Fprintf(w, "`key` is: %s\n", key)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at: 8080")
	http.ListenAndServe(":8080", nil)
}

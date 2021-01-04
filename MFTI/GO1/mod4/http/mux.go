package main

import (
	"fmt"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "serve Mux")
}

func main() {
	// simple http.HandleFunc use deafult global multiplex
	// but we can use custom mux
	mux := http.NewServeMux() // create mux
	mux.HandleFunc("/", rootHandler)

	// customization
	server := http.Server{
		Addr:         ":8080", // port
		Handler:      mux,     // mux
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

// note: we can create few different mux to different handling

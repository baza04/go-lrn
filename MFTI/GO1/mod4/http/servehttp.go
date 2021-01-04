package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Name string
}

// ServeHTTP is method which can handle request for few URL's
// and work with structs and other types
func (h *Handler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	fmt.Fprintf(w, "Handler name: %s, URL: %s", h.Name, r.URL.String())
}

func main() {
	testHandler := Handler{Name: "test"}
	http.Handle("/test/", testHandler)

	rootHandler := Handler{Name: "root"}
	http.Handle("root", rootHanle)

	fmt.Println("Listening port: 8080")
	http.ListenAndServe(":8080", nil)
}

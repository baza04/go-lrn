package main

import (
	"fmt"
	"net/http"
)

w.Header().Set("Content-Type", "text/html")
w.Write([]byte(`
<img src="/static/img/gopher.png" />
`))
}

func main() {
mux := http.NewServeMux()
mux.HandleFunc("/", handler)

// static files parser(handler) first args like link to next calling
// http.Dir  read static files
static := http.StripPrefix("/data/",
	http.FileServer(http.Dir("./static")),
)
// parse static files
http.Handle("/data/", static)
server := http.Server{
	Addr:    ":8080",
	Handler: mux,
}
fmt.Println("starting server at :8080")
server.ListenAndServe()
}

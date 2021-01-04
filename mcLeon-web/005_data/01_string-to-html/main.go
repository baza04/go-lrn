package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	port := os.Args[1]

	arg := []int{13, 543, 63}
	// fmt.Printf("Listen port: %i", port)
	// err = tpl.ExecuteTemplate(indexHandler, "index.html", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", arg)
	if err != nil {
		log.Fatalln(err)
	}
	http.ListenAndServe(":"+port, nil)
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	err := tpl.ExecuteTemplate("tpl.gohtml", 42)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

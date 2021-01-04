package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "no one can`t be forever")
	if err != nil {
		log.Fatalln(err)
	}
}

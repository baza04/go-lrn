package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	argsMap := map[string]string{
		"yes": "no",
	}
	err := tpl.Execute(os.Stdout, argsMap)
	if err != nil {
		log.Fatalln(err)
	}
}

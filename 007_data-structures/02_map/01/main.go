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
	// sages := []string{"Cris", "Turoph", "Sam", "Mike"}

	Map := map[string]string{
		"Sam":    "Smith",
		"Tom":    "Cruz",
		"Jesica": "Alba",
	}
	err := tpl.Execute(os.Stdout, Map)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Auth data
type Auth struct {
	Login string
	Pass  string
	Email string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	a := Auth{
		Login: "aaa",
		Pass:  "a1a1a1",
		Email: "aaa@testmail.com",
	}

	b := Auth{
		Login: "bbb",
		Pass:  "b2b2b2",
		Email: "bbb@testmail.com",
	}

	c := Auth{
		Login: "ccc",
		Pass:  "c3c3c3",
		Email: "ccc@testmail.com",
	}

	SOS := []Auth{a, b, c}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", SOS)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("is the programm stil work?")
}

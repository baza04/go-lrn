package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type Auth struct {
	Fname string
	Sname string
	Login string
	Email string
}

type car struct {
	Manifacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []Auth
	Transport []car
}

var tpl *template.Template
var tpl1 *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	a := Auth{
		Fname: "Avram",
		Sname: "Linkoln",
		Login: "LinkA",
		Email: "Alink@testmial.com",
	}

	b := Auth{
		Fname: "Theodor",
		Sname: "Ruspheld",
		Login: "RusTh@testmail.com",
		Email: "Trusp@testmail.com",
	}

	c := Auth{
		Fname: "John",
		Sname: "Kenedy",
		Login: "KeyJo",
		Email: "Joke@testmail.com",
	}

	d := car{
		Manifacturer: "Toyota",
		Model:        "Tundra",
		Doors:        5,
	}

	e := car{
		Manifacturer: "Nissan",
		Model:        "GTX",
		Doors:        3,
	}

	f := car{
		Manifacturer: "Mitsubishi",
		Model:        "Eclipse",
		Doors:        3,
	}

	ad := []Auth{a, b, c}
	cd := []car{d, e, f}

	Data := items{
		Wisdom:    ad,
		Transport: cd,
	}
	/* err := tpl.ExecuteTemplate(os.Stdout, "newTpl", nil)
	if err != nil {
		log.Fatalln(err)
	} */
	fmt.Println("test")
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Data)
	if err != nil {
		log.Fatalln(err)
	}
}

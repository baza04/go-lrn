package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manifacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	a := sage{
		Name:  "Abai",
		Motto: `"Don't be pround if you haven't known a wisdom"`,
	}

	b := sage{
		Name:  "Ybyray",
		Motto: `"Children let's go  to study"`,
	}

	c := sage{
		Name:  "Saken",
		Motto: `"Meaning of some words is growth with time..."`,
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

	sages := []sage{a, b, c}
	cars := []car{d, e, f}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

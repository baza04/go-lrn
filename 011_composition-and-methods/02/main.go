package main

import (
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1, p2 := doubleZero{
		person{
			Name: "Amir",
			Age:  25,
		},
		false,
	}, doubleZero{
		person{
			Name: "Clare",
			Age:  25,
		},
		true,
	}
	arr := []doubleZero{p1, p2}
	tpl.Execute(os.Stdout, arr)
}

package main

import (
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1, p2 := person{
		Name: "Amir",
		Age:  25,
	}, person{
		Name: "Clare",
		Age:  25,
	}
	arr := []person{p1,p2}
	tpl.Execute(os.Stdout, arr)
}

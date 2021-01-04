package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Arg struct {
	FName string
	SName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	Make := Arg{
		FName: "Amir",
		SName: "Bazarbay",
	}
	// slice := []int{23, 25, 26, 28}

	Den := Arg{
		FName: "Daniyar",
		SName: "Dyusimbayev",
	}

	Devstack := Arg{
		FName: "Bekzhan",
		SName: "Sattarkulov",
	}
	sOS := []Arg{Make, Den, Devstack}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sOS)
	if err != nil {
		log.Fatalln(err)
	}

}

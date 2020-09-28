package main

import (
	"os"
	"text/template"
)

type semester struct {
	Term    string
	Courses []course
}

type course struct {
	Number, Name, Units string
}


type year struct {
	Fail, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	s1 := year{
		Fail: semester{
			Term: "Fail",
			Courses: []course{
				course{"numberOne", "Computer Sciense", "3"},
				course{"numberTwo", "Data sciense", "6"},
				course{"numberThree", "Algo Basics", "5"},
			},
		},
		Spring: semester{
			Term: "Bim",
			Courses: []course{
				course{"num4","DS", "4"},
				course{"num5","CS", "7"},
				course{"num6","AB", "4"},
			},
		},
		Summer: semester{
			Term: "Bim",
			Courses: []course{
				course{"nb7", "ds", "5"},
				course{"nb8", "cs", "7"},
				course{"nb9", "ab", "8"},
			},

		},
	}

	tpl.Execute(os.Stdout, s1)
}

package main

import "fmt"

type hotdoc int

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fName, p.lName, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fName, sa.lName, `says, "Shaken, not stirred."`)
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	saySomething(sa1)
}

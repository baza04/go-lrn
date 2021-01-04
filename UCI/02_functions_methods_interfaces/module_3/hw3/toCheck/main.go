package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}
func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}


func main() {
	var name, info string
	var anim Animal 
	cow := Animal{food: "grass",
				  locomotion: "walk",
				  noise: "moo"}
	bird := Animal{food: "worms",
				  locomotion: "fly",
				  noise: "peep"}
	snake := Animal{food: "mice",
				  locomotion: "slither",
				  noise: "hsss"}
	
	for {
		fmt.Print("> ")
		fmt.Scan(&name)
		fmt.Scan(&info)
		if name=="cow"{
			anim = cow
		} else if name=="bird"{
			anim = bird
		} else if name=="snake"{
			anim = snake
		} else {
			anim = Animal{food: "", locomotion: "", noise: ""} // For incorrect animal name, the program prints a blank line
		}

		if info=="eat" {
			anim.Eat()
		} else if info=="move" {
			anim.Move()
		} else if info=="speak" {
			anim.Speak()
		}
	}
}
	
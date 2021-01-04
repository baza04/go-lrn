package main

import (
	"fmt"
	"strings"
)

// Animal info
type Animal struct {
	food, motionMethod, sound string
}

type request struct {
	name, info string
}

func main() {
	printDesc()
	data := map[string]Animal{}
	initInfoArr(data)

	for {
		r := request{}
		if !r.readReq() {
			fmt.Println("!!! Incorrect name or requested info, please try again to type 'animal name', then 'info type'")
			continue
		}
		temp := data[r.name]
		temp.printInfo(r.name, r.info)
	}
}

func (r *request) readReq() bool {
	fmt.Print("> ")
	fmt.Scan(&r.name)
	if !r.isNameValid() {
		return false
	}
	fmt.Print("> ")
	fmt.Scan(&r.info)
	if !r.isReqValid() {
		return false
	}
	return true
}

func (r *request) isNameValid() bool {
	r.name = strings.TrimSpace(r.name)
	return r.name == "cow" || r.name == "snake" || r.name == "bird"
}

func (r *request) isReqValid() bool {
	r.info = strings.TrimSpace(r.info)
	return r.info == "food" || r.info == "motion" || r.info == "sound"
}

func initInfoArr(arr map[string]Animal) {
	arr["cow"] = Animal{"grass", "walk", "moo"}
	arr["bird"] = Animal{"worms", "fly", "peep"}
	arr["snake"] = Animal{"mice", "slither", "hsss "}
}

func (a *Animal) printInfo(name, info string) {
	fmt.Print(name)
	switch info {
	case "food":
		a.Eat()
	case "motion":
		a.Move()
	case "sound":
		a.Speak()
	}
}

// Eat print what choosen animal eat
func (a *Animal) Eat() {
	fmt.Printf(" eat: %s\n", a.food)
}

// Move print how choosen animal move
func (a *Animal) Move() {
	fmt.Printf(" move by: %s\n", a.motionMethod)
}

// Speak print how choosen animal speak
func (a *Animal) Speak() {
	fmt.Printf(" speak by %s\n", a.sound)
}

func printDesc() {
	fmt.Println(`		*** DESCRIPTION ***
This programm print info about animals: cow, bird, snake
you can know what this animals eat, how they are speak and move.

To get info please type
> animal name like: "cow" or "bird" or "snake"
> info you want to know like: "food" or "motion" or "sound"
Let's Start!`)
}

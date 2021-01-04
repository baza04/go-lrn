package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal can print options of animals
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct to realize Eat(), Move(), Speak() for cow type
type Cow struct{}

// Bird struct to realize Eat(), Move(), Speak() for bird type
type Bird struct{}

// Snake struct to realize Eat(), Move(), Speak() for snake type
type Snake struct{}

func (a *Cow) Eat()   { fmt.Println("grass") }
func (a *Cow) Move()  { fmt.Println("walk") }
func (a *Cow) Speak() { fmt.Println("moo") }

func (a *Bird) Eat()   { fmt.Println("worms") }
func (a *Bird) Move()  { fmt.Println("fly") }
func (a *Bird) Speak() { fmt.Println("poop") }

func (a *Snake) Eat()   { fmt.Println("mice") }
func (a *Snake) Move()  { fmt.Println("slither") }
func (a *Snake) Speak() { fmt.Println("hsss") }

// var list map[string]Animal

func main() {
	printDesc()
	list := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter request please:\n> ")
		req, _ := reader.ReadString('\n')
		handleRequests(req, list)
	}
}

func handleRequests(req string, list map[string]Animal) {
	args := strings.Fields(req)
	fmt.Println(args)

	if args[0] == "newanimal" {
		switch args[2] {
		case "cow":
			list[args[1]] = new(Cow)
		case "bird":
			list[args[1]] = new(Bird)
		case "snake":
			list[args[1]] = new(Snake)
		}
		fmt.Println("Created it!")
	}
	if args[0] == "query" {
		if animal, ok := list[args[1]]; ok {
			switch args[2] {
			case "food":
				animal.Eat()
			case "motion":
				animal.Move()
			case "sound":
				animal.Speak()
			}
		} else {
			fmt.Println("Not found!")
		}
	}
}

func printDesc() {
	fmt.Println(`			*** DESCRIPTION ***

This program can get info about exist or create new animal for next types of animals:
- cow
- bird 
- snake

To create new animal type:
- newanimal "animal name" "animal type" (cow or bird or snake)

To get info please type:
 - query "animal name"  "info" (food or motion or sound)

Let's Start!`)
}

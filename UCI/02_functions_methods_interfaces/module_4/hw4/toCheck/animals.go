package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {}
type Bird struct {}
type Snake struct {}


func (cow Cow) Eat() {
    fmt.Println("grass")
}


func (bird Bird) Eat() {
    fmt.Println("worms")
}


func (snake Snake) Eat() {
    fmt.Println("mice")
}


func (cow Cow) Move() {
    fmt.Println("walk")
}


func (bird Bird) Move() {
    fmt.Println("fly")
}


func (snake Snake) Move() {
    fmt.Println("slither")
}


func (cow Cow) Speak() {
    fmt.Println("moo")
}


func (bird Bird) Speak() {
    fmt.Println("peep")
}


func (snake Snake) Speak() {
    fmt.Println("hsss")
}


func main() {
    var animals = make(map[string] Animal, 0)
    var species = map[string] Animal {
        "cow":  new(Cow),
        "bird": new(Bird),
        "snake": new(Snake),
    }
    var methods = map[string] func(Animal) {
        "eat": Animal.Eat,
        "move": Animal.Move,
        "speak": Animal.Speak,
    }
    in := bufio.NewScanner(os.Stdin)
    
    fmt.Println("Welcome!")
    
    for {
        fmt.Println("> Type (newanimal name type)/(query name method): ")
        in.Scan()
        inputs := strings.Split(in.Text(), " ")
        
        if inputs[0] == "newanimal" {
            animals[inputs[1]] = species[inputs[2]]
        } else if inputs[0] == "query" {
            methods[inputs[2]](animals[inputs[1]])
        } else {
            fmt.Println("Option not recognized")
        }
    }
}
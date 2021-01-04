package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type person struct {
	fname, lname []rune
}

func main() {
	fmt.Println("Enter path to file:")
	var path string
	fmt.Scan(&path)
	slice := []person{}

	data, _ := ioutil.ReadFile(path)
	arr := strings.Fields(string(data))
	
	var fname, lname []rune
	for index, value := range arr {
		if index%2 == 0 {
			fname = []rune(value)
		} else {
			lname = []rune(value)
			slice = append(slice, person{fname, lname})
		}
	}
	for _, obj := range slice {
		fmt.Println(string(obj.fname), string(obj.lname))
	}
}

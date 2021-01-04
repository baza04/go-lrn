package main

import "fmt"

func main() {
	pi := 3.14
	fmt.Println("Enter some float number like Pi:", pi)
	var input float64
	fmt.Scan(&input)
	fmt.Println(int(input))
}

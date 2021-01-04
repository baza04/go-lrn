package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := map[string]string{}
	m["name"] = readInput("Would you kindly type your name:")
	m["address"] = readInput("Please enter your address:")

	fmt.Printf("JSON representation: %s\n", makeJSON(m))
}

func readInput(text string) string {
	var input string
	fmt.Println(text)

	nb, err := fmt.Scan(&input)
	if nb < 1 || nb > 20 || err != nil {
		fmt.Println("!!! Incorrect input !!!\nTry to run app again")
		os.Exit(1)
	}
	return input
}

func makeJSON(data map[string]string) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return json
}

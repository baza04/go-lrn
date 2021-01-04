package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	name = strings.Replace(name, "\n", "", -1)

	fmt.Print("Enter address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	address = strings.Replace(address, "\n", "", -1)

	personMap := map[string]string{"name": name, "address": address}

	byteArr, err := json.Marshal(personMap)
	fmt.Println(byteArr)
	fmt.Println(string(byteArr))
}

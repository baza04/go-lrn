package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Would you kindly to enter some string:")

	for {
		findian()
	}
}

func findian() {
	fmt.Println("**************************************")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str = strings.ToLower(str[:len(str)-1]) // str[len(str)-1] need to cut "\n"
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	fmt.Println("\nEnter one more string please:")
}

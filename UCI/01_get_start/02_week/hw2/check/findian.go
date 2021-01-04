package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your string: ")

	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.Replace(text, "\n", "", 1))

	isI := strings.Index(text, "i")
	isA := strings.Index(text, "a")
	isN := strings.Index(text, "n")

	isILetterPresented := isI == 0
	isALetterPresented := isILetterPresented && isA > 0
	isNLetterPresented := isILetterPresented && isN > isA

	if isILetterPresented && isALetterPresented && isNLetterPresented {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

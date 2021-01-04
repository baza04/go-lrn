package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"io"
)

func main(){
	type Name struct{
		FirstName string
		LastName string
	}
	fileContent := make([]Name, 0, 20)
	var fileName string

	fmt.Print("Insert the filename to read in the current directory: ")
	fmt.Scan(&fileName)
	fmt.Printf("\n")
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to read the file "+fileName)
	}
	reader := bufio.NewReader(f)
	fmt.Println("Reading file: "+fileName)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF{
			defer f.Close()
			break
		}
		if line == nil {
			fmt.Println("Line is empty")
			defer f.Close()
			return
		}
		ll := string(line)

		fileContent = append(fileContent, Name{
			FirstName: strings.Split(ll, " ")[0],
			LastName: strings.Split(ll, " ")[1],
		})
	}
	defer f.Close()
	fmt.Println("Iterating over the slice")
	for idx, c := range fileContent{
		fmt.Printf("ROW: %d - First Name: %s, Last Name: %s \n", idx+1, c.FirstName, c.LastName)
	}
}
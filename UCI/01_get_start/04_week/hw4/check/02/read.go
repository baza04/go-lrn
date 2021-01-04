package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var textfile_name string
	fmt.Println("Enter text file name: ")
	fmt.Scan(&textfile_name)
	fmt.Println("Reading from file ", textfile_name)
	//open file
	// f, _ := os.Open("C:\\Users\\mscra\\Desktop\\KB\\Coursera_Golang\\hello\\input.txt")
	f, _ := os.Open(textfile_name)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var text []string
	//reading each line from file
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	//closing file
	f.Close()

	slice := make([]name, 1)
	i := 0
	for _, each_line := range text {
		//splitting contents in line into two parts
		s := strings.Fields(each_line)
		if i == 0 {
			slice[0].fname = s[0]
			slice[0].lname = s[1]
			i++
		} else {
			newItem := name{fname: s[0], lname: s[1]}
			slice = append(slice, newItem)
		}

	}

	for _, each_slice := range slice {
		fmt.Println("First name : ", each_slice.fname, " , last name : ", each_slice.lname)
	}
}

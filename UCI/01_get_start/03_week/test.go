package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// func main() {
// 	// x := [4]int{0, 1, 2, 4}
// 	y := make([]int, 4, 9)
// 	y = append(y, 1, 8, 8, 8, 8)
// 	fmt.Println(len(y), cap(y), y)
// 	fmt.Println(reflect.TypeOf(y))
// }

// ( "fmt" "io/ioutil" "strings" ) //constant size of name field const SIZE=20;
//def for Name struct
type Name struct {
	fname [20]byte
	lname [20]byte
}

func main() {
	var fileName string
	// define slice of type name
	names := make([]Name, 0)
	fmt.Println("Enter file name without .txt extension")
	fmt.Scan(&fileName)

	//reading data from file
	data, err := ioutil.ReadFile(fileName + ".txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
	}

	//iterate though data
	for _, lineData := range strings.Split(string(data), "\n") {
		nameList := strings.Fields(lineData)
		if len(nameList) > 0 {
			fmt.Println(nameList[0], nameList[1])
			names = append(names, Name{cnvrt(nameList[0]), cnvrt(nameList[1])})
		}
	}
	
	// iterate through slice
	for i, val := range names {
		fmt.Printf("%d. Name : %s %s\n", i+1, val.fname, val.lname)
	}
}

//function to convert string to byte array
func cnvrt(str string) [20]byte {
	var b [20]byte
	copy(b[:], []byte(str[:len(str)]))
	return b
}

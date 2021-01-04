// go build gen/* && ./codegen.exe pack/packer.go  pack/marshaller.go
package main

import "fmt"

// lets generate code for this struct
// cgen: binpack
type User struct {
	ID       int
	RealName string `unpack: "-"`
	Login    string
	Flags    int
}
type Avatar struct {
	ID  int
	Url string
}

var test = 42

func main() {
	// data in binary representation
	data := []byte{
		128, 36, 17, 0, // int

		9, 0, 0, 0, // line len
		118, 46, 114, 111, 109, 97, 110, 111, 118, // str in binary

		16, 0, 0, 0, // some int
	}

	u := new(User)
	u.Unpack(data)
	fmt.Printf("Unpacked user %#v", u)
}

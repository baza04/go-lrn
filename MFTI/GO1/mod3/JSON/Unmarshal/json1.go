package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	phone    string
}

var jsonStr = `{"id":42, "username": "Make", "phone": "132"}`

func main() {
	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("struct: \n\t%#v\n\n", u)

	u.phone = "978163"
	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string: \n\t%s\n", string(result))
}

// struct is template for json object
// if some field of struct named in lower case
// it will give permission only for funcs or methods in this pkg
// and pkg "encoding/json" haven`t permisson for this field

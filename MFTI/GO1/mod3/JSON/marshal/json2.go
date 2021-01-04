package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	ID       int `json:"user_id,string"`
	Username string
	Address  string `json:",omitempty"`
	Company  string `json:"-"`
}

// Marshal json object with customized fields named
func main() {
	u := &User1{
		ID:       4,
		Username: "Make",
		Address:  "050000",
		Company:  "01 alem",
	}

	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string:\t%s\n", string(result))
}

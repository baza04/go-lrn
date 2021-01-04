package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr1 = `[
	{"id":17, "username": "iivan", "phone": 0},
	{"id":"18", "address": "none", "company": "alem"}	
]`

func main() {
	// Unmarshal from JSON with unknown fields
	data := []byte(jsonStr1)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("unpacked in empty interface:\n%#v\n\n", user1)

	// marshal map of interface to JSON
	user2 := map[string]interface{}{
		"id":       42,
		"username": "make",
	}

	var user2i interface{} = user2
	result, err := json.Marshal(user2i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON from inteface: %s\n", result)
}

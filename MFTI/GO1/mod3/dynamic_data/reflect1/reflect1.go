package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int
	RealName string `unpack: "-"`
	Login    string
	Flags    int
}

func main() {
	u := &User{
		ID:       42,
		RealName: "Make",
		Flags:    32,
	}

	err := PrintReflect(u)
	if err != nil {
		panic(err)
	}

}

func PrintReflect(u *User) error {
	// get struct info from recieved object by reflect pkg
	val := reflect.ValueOf(u).Elem()

	// get num of struct fields
	fmt.Printf("%T have %d fields: \n", u, val.NumField())

	// start loop by struct fields
	for i := 0; i < val.NumField(); i++ {
		// get value of choosen field
		valueField := val.Field(i)
		// get info of choosen field
		typeField := val.Type().Field(i)

		// print name, data type, value and tag of choosen field if exist
		fmt.Printf("\tname=%v, type=%v, value=%v, tag='%v'\n", typeField.Name, typeField.Type.Kind(), valueField, typeField.Tag)
	}
	return nil
}

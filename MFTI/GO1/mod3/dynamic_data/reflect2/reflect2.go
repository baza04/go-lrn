package main

import (
	"bytes"
	"encoding/binary"
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
	// data in binary representation
	data := []byte{
		128, 36, 17, 0, // int

		9, 0, 0, 0, // line len
		118, 46, 114, 111, 109, 97, 110, 111, 118, // str in binary

		16, 0, 0, 0, // some int
	}

	u := new(User)
	err := UnpackReflect(u, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", u)
}

func UnpackReflect(u interface{}, data []byte) error {
	r := bytes.NewReader(data) // read from binary data to buffer

	// get struct from object
	val := reflect.ValueOf(u).Elem()

	for i := 0; i < val.NumField(); i++ {
		// get value of current field in struct User (which init  above)
		valueField := val.Field(i)
		// get type of current field in of struct User
		typeField := val.Type().Field(i)

		if typeField.Tag.Get("unpack") == "-" {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			var value uint32 // init temp var

			// read from buffer (r) write to temp var
			binary.Read(r, binary.LittleEndian, &value)
			// set data type to value in temp var
			valueField.Set(reflect.ValueOf(int(value)))
		case reflect.String:
			var lenRaw uint32                            // init var to string len
			binary.Read(r, binary.LittleEndian, &lenRaw) // write len

			dataRaw := make([]byte, lenRaw)               // init slice of byte for string value
			binary.Read(r, binary.LittleEndian, &dataRaw) // write value of string to temp var
			valueField.SetString(string(dataRaw))         // set data type
		default:
			return fmt.Errorf("bad type: %v for field %v", typeField.Type.Kind(), typeField.Name)
		}
	}
	return nil
}

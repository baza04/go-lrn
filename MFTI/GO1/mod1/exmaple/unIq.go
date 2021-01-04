package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	fmt.Println("Start scanning:")
	pref := ""
	for in.Scan() {
		text := in.Text()

		if text < pref {
			return fmt.Errorf("file not sorted")
		} else if text == pref {
			continue
		}
		pref = text
		fmt.Fprintln(output, text)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

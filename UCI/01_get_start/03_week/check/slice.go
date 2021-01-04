package main

import (
	b "bufio"
	f "fmt"
	o "os"
	so "sort"
	sc "strconv"
	s "strings"
)

func checkErr(err error) {
	if err != nil {
		f.Println("You must enter an Integer value")
	}
}

func main() {
	replacer := s.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
	reader := b.NewReader(o.Stdin)
	slc := make([]int, 3)

	f.Println("Slice Sorter")
	f.Println("======================")
	i := 0
	for {
		f.Print("Enter integer: ")
		t, _ := reader.ReadString('\n')
		t = replacer.Replace(t)
		if s.ToUpper(t) == "X" {
			break
		}

		num, err := sc.Atoi(t)
		checkErr(err)
		if err == nil {
			if i < len(slc)-1 {
				slc[i] = num
			} else {
				slc = append(slc, num)
			}
			so.Ints(slc)
			f.Println(slc)
		}
		i++
	}
}

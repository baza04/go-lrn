package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func sorting(id int, sli []int, c chan []int) {
	sort.Ints(sli)
	fmt.Println("Goroutine #", id, goid(), "Array sorted -> ", sli)
	c <- sli
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	//channel
	c := make(chan []int)

	//scanner
	scanner := bufio.NewScanner(os.Stdin)
	for {
		//sli integers
		sliInt := make([]int, 0, 100)
		fmt.Print("enter serie of integer: ")

		//Reading
		scanner.Scan()
		command := scanner.Text()
		command = strings.ToLower(command)
		commands := strings.Fields(command)
		for i := 0; i < len(commands); i++ {
			textInt, _ := strconv.Atoi(commands[i])
			sliInt = append(sliInt, textInt)
		}

		fmt.Println("Array of integers -> ", sliInt)
		//spliting Array
		const interval int = 4
		var counter int = 0
		lenSliInt := len(sliInt)
		if lenSliInt >= interval {
			limit := int(math.Round(float64(lenSliInt) / float64(interval)))
			fmt.Println(limit)
			//Executing goroutines
			for i := 0; i < lenSliInt; i += limit {
				batch := sliInt[i:min(i+limit, lenSliInt)]
				// batch := sliInt[i : i+limit]
				fmt.Println(batch)
				go sorting(counter, batch, c)
				counter++
			}
			//new array from merged arrays
			mergedSliInt := make([]int, 0)
			//Consuming channel
			for i := 0; i < counter; i++ {
				sortedSli := <-c
				mergedSliInt = append(mergedSliInt, sortedSli...)
			}
			//Sort merge sli
			sort.Ints(mergedSliInt)
			fmt.Println("Array merged -> ", mergedSliInt)
		}

	}
	fmt.Println("End!")
}

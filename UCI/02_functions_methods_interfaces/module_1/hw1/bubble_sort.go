package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Would you kindly to type up to 10 integers?")
	arr := strToInt(readInt())
	arr = bubbleSort(arr)
	fmt.Printf("sorted array: %v\n", arr)
}

func readInt() []string {
	replacer := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = replacer.Replace(input)

	return strings.Fields(input)
}

func strToInt(strArr []string) (result []int) {
	for _, value := range strArr {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return
}

func bubbleSort(arr []int) []int {
	lnA := len(arr)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < lnA-1; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		lnA--
	}

	return arr
}

func swap(arr []int, index int) {
	(arr)[index+1], (arr)[index] = (arr)[index], (arr)[index+1]
}

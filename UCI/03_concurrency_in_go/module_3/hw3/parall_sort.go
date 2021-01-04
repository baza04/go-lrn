package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("input from 8 to 20 integers without commas (just integers)\nExample: 253 4 76 30 64 19 7 84 163 69 71 82 57 22\n\nEnter numbers please\n> ")
	strArr := reader()
	arr := strToInt(strArr)
	sortedArr := toSort(arr)
	fmt.Println(sortedArr)
}

// just read input & return arr
func reader() []string {
	replacer := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	if !isValid(input) {
		fmt.Printf("Incorrect Input!!! Please try again\n> ")
		return reader()
	}
	return strings.Fields(replacer.Replace(input))
}

func isValid(input string) bool {
	for _, rune := range input {
		if !unicode.IsDigit(rune) && !unicode.IsSpace(rune) {
			return false
		}
	}
	return true
}

// convert str arr to int arr
func strToInt(strArr []string) (rArr []int) {
	for _, v := range strArr {
		num, err := strconv.Atoi(v)
		if err == nil {
			rArr = append(rArr, num)
		}
	}
	return
}

// divide total arr to 4 parts and put it to mergeSort by goroutines
// then wait when some of 4 sort will done then merge to result
func toSort(arr []int) (sortedArr []int) {
	quarter, start := len(arr)/4, 0
	c := make(chan []int, 4)
	for i := 0; i < len(arr); i++ {
		if isQuarter(i, quarter, len(arr)) {
			go toParral(arr[start:start+quarter], c)
			start += quarter
		}
	}
	go toParral(arr[start:], c)
	wg.Wait()

	for i := 0; i < 4; i++ {
		sortedArr = append(sortedArr, <-c...) // get sorted quarter arr from channel
		sortedArr = mergeSort(sortedArr)
	}
	return
}

// just check to valid quarter
func isQuarter(i, quarter, len int) bool {
	return i != 0 && i%quarter == 0 && len-i >= quarter
}

// add WaitGroup send sorted part to channel
func toParral(arr []int, c chan []int) {
	wg.Add(1)
	sorted := mergeSort(arr)
	c <- sorted
	fmt.Println(arr)
	wg.Done()
}

// not comment
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	left := mergeSort(arr[middle:])
	right := mergeSort(arr[:middle])

	return merge(left, right)
}

func merge(l, r []int) []int {
	result := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(result, r...)
		}
		if len(r) == 0 {
			return append(result, l...)
		}

		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	return result
}

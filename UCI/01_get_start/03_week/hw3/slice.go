package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := []int{}
	ok := false
	fmt.Printf("Would you kindly to type some integer?\ninput: ")

	for {
		slice, ok = readInt(slice)
		if ok {
			fmt.Printf("slice: %v\n\nType one more please:\ninput: ", slice)
		}
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

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

func readInt(arr []int) ([]int, bool) {
	var input string
	fmt.Scan(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("!!! Incorrect input: \"%s\" !!!\n\nEnter correct interger please:\ninput: ", input)
		return arr, false
	}
	arr = append(arr, num)
	return mergeSort(arr), true
}

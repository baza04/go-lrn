package main

import "fmt"

func swap(swapSliceEle []int, index int) {
	swapSliceEle[index], swapSliceEle[index+1] = swapSliceEle[index+1], swapSliceEle[index]
}

func BubbleSort(sliceIpOp []int) {
	var index int
	lengthOfSlice := len(sliceIpOp)
	fmt.Println("length of slice: ", lengthOfSlice)
	for index < lengthOfSlice-1 {
		index2 := 0
		for index2 < lengthOfSlice-index-1 {
			if sliceIpOp[index2] > sliceIpOp[index2+1] {
				swap(sliceIpOp, index2)
			}
			index2++
		}
		index++
	}
}

func main() {
	sliceInt := make([]int, 10)
	limitInt := 0
	fmt.Printf("Enter elements : ")
	for limitInt < 10 {
		fmt.Scan(&sliceInt[limitInt])
		limitInt++
	}
	fmt.Println("Slice before bubble sort: ", sliceInt)
	BubbleSort(sliceInt)
	fmt.Println("Slice after bubble sort: ", sliceInt)
}

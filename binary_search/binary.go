package main

import (
	"fmt"
)

func binarySearch(array []int, target int, start int, end int) int {
	middle := int((end + start) / 2)

	if end < start {
		return -1
	}

	if array[middle] > target {
		return binarySearch(array, target, start, middle)
	} else if array[middle] > target {
		return binarySearch(array, target, start, middle)
	} else if array[middle] < target {
		return binarySearch(array, target, middle+1, end)
	} else /*target == array[middle]*/ {
		return middle
	}
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10

	fmt.Println(binarySearch(array, target, 0, len(array)-1))
}

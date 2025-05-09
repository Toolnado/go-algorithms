package main

import (
	"fmt"
	"slices"
)

func main() {
	elements := []int{15, 48, 26, 18, 41, 86, 29, 51, 20}
	slices.Sort(elements)
	fmt.Println(BinarySearch(elements, 48))
}

func BinarySearch(elements []int, findElement int) int {
	low, high := 0, len(elements)-1
	for low <= high {
		mid := low + (high-low)/2
		switch {
		case elements[mid] == findElement:
			return mid
		case findElement < elements[mid]:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return -1
}

package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	InsertionSorter(nums)
	fmt.Println("sort nums:", nums)
}

func InsertionSorter(elements []int) {
	for i := 1; i < len(elements); i++ {
		j := i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j--
		}
	}
}

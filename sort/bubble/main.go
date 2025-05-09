package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	BubbleSorter(nums)
	fmt.Println("sort nums:", nums)
}

func BubbleSorter(elements []int) {
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for i := 1; i < len(elements); i++ {
			if elements[i-1] > elements[i] {
				elements[i-1], elements[i] = elements[i], elements[i-1]
				isSwapped = true
			}
		}
	}
}

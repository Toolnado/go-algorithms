package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	SelectionSorter(nums)
	fmt.Println("sort nums:", nums)
}

func SelectionSorter(elements []int) {
	for i := 0; i < len(elements)-1; i++ {
		min := i
		for j := i + 1; j <= len(elements)-1; j++ {
			if elements[j] < elements[min] {
				min = j

			}
		}
		elements[min], elements[i] = elements[i], elements[min]
	}
}

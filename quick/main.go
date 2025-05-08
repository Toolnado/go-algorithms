package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	QuickSorter(nums, 0, len(nums)-1)
	fmt.Println("sort nums:", nums)
}
func QuickSorter(elements []int, below int, upper int) {
	if below < upper {
		part := divide(elements, below, upper)
		QuickSorter(elements, below, part-1)
		QuickSorter(elements, part+1, upper)
	}
}

func divide(elements []int, below int, upper int) int {
	i := below
	for j := below; j < upper; j++ {
		if elements[j] <= elements[upper] {
			elements[i], elements[j] = elements[j], elements[i]
			i++
		}
	}
	elements[i], elements[upper] = elements[upper], elements[i]
	return i
}

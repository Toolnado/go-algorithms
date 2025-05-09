package main

import "fmt"

func main() {
	elements := []int{15, 48, 26, 18, 41, 86, 29, 51, 20}
	fmt.Println(LinearSearch(elements, 48))
}

func LinearSearch(elements []int, findElement int) int {
	for i, e := range elements {
		if e == findElement {
			return i
		}
	}
	return -1
}

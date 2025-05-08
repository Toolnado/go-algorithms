package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	BubbleSorter(nums)
	fmt.Println("sort nums:", nums)
}

func BubbleSorter(nums []int) {
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				isSwapped = true
			}
		}
	}
}

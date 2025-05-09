package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	fmt.Println("sort nums:", MergeSorter(nums))
}

func MergeSorter(elements []int) []int {
	if len(elements) < 2 {
		return elements
	}

	middle := len(elements) / 2

	return join(MergeSorter(elements[:middle]), MergeSorter(elements[middle:]))
}

func join(leftArr, rightArr []int) []int {
	length := len(leftArr) + len(rightArr)
	arr := make([]int, length, length)
	lp, rp := 0, 0

	for i := 0; i < length; i++ {
		if lp >= len(leftArr) && rp < len(rightArr) {
			arr[i] = rightArr[rp]
			rp++
		} else if rp >= len(rightArr) && lp < len(leftArr) {
			arr[i] = leftArr[lp]
			lp++
		} else if leftArr[lp] < rightArr[rp] {
			arr[i] = leftArr[lp]
			lp++
		} else {
			arr[i] = rightArr[rp]
			rp++
		}
	}

	return arr
}

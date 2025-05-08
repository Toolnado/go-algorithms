package main

import "fmt"

func main() {
	nums := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("nums:", nums)
	ShellSorter(nums)
	fmt.Println("sort nums:", nums)
}

func ShellSorter(elements []int) {
	var (
		n         = len(elements)
		intervals = []int{1}
		k         = 1
	)

	for {
		var interval int
		interval = power(2, k) + 1
		if interval > n-1 {
			break
		}
		intervals = append([]int{interval}, intervals...)
		k++
	}

	var interval int
	for _, interval = range intervals {
		var i int
		for i = interval; i < n; i += interval {
			var j int
			j = i
			for j > 0 {
				if elements[j-interval] > elements[j] {
					elements[j-interval], elements[j] = elements[j], elements[j-interval]
				}
				j = j - interval
			}
		}
	}
}

func power(exponent int, index int) int {
	var power int
	power = 1
	for index > 0 {
		if index&1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}

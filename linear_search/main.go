package main

import "fmt"

func LinearSearch(nums []int, value int) int {
	for i := range nums {
		if nums[i] == value {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{2, 4, 3, 1, 5, 7, 6}
	fmt.Println(LinearSearch(nums, 5))
}

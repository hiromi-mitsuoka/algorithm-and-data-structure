package main

import "fmt"

// func BinarySearch(nums []int, value int) int {
// 	max := len(nums)
// 	min := 0
// 	midian := (max - min) / 2
// 	// fmt.Println(min, max, midian)

// 	for max >= min {
// 		if nums[midian] < value {
// 			min = midian + 1
// 			midian = (max + min) / 2
// 			// fmt.Println(min, max, midian)
// 		} else if nums[midian] > value {
// 			max = midian - 1
// 			midian = (max + min) / 2
// 			// fmt.Println(min, max, midian)
// 		} else {
// 			return midian
// 		}
// 	}
// 	return -1
// }

// Use recursive
func BinarySearch(nums []int, value int, min int, max int, midian int) int {
	fmt.Println(min, max, midian)
	for max >= min {
		if nums[midian] < value {
			min = midian + 1
			midian = (max + min) / 2
			return BinarySearch(nums, value, min, max, midian)
		} else if nums[midian] > value {
			max = midian - 1
			midian = (max + min) / 2
			return BinarySearch(nums, value, min, max, midian)
		} else {
			return midian
		}
	}
	return -1
}

func main() {
	// nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(BinarySearch(nums, 7))

	// Use recursive
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	max := len(nums) - 1
	min := 0
	midian := (max - min) / 2
	fmt.Println(BinarySearch(nums, 6, min, max, midian))
}

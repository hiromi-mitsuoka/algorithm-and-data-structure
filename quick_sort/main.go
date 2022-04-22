package main

import (
	"fmt"
	"sort"
)

func partition(nums []int, low int, high int) int {
	i := low - 1
	pivot := nums[high]
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func QuickSort(nums []int, low int, high int) []int {
	if low < high {
		partition_index := partition(nums, low, high)
		QuickSort(nums, low, partition_index-1)
		QuickSort(nums, partition_index+1, high)
	}
	return nums
}

func main() {
	nums := []int{2, 4, 3, 1, 5, 7, 6}
	low := 0
	high := len(nums) - 1
	fmt.Println(QuickSort(nums, low, high))

	// Use package
	sort.Ints(nums)
	fmt.Println(nums)
}

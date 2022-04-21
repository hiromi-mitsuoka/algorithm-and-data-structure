package main

import "fmt"

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			// fmt.Println(i, j, nums[j], nums[j+1])
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{2, 4, 3, 1, 5, 7, 6}
	fmt.Println(BubbleSort(nums))
}

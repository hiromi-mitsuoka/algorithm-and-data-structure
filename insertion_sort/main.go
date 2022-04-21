package main

import "fmt"

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			j := i
			tmp := nums[i]
			for 0 < j && nums[j-1] > tmp {
				nums[j] = nums[j-1]
				j--
				fmt.Printf("i = %d, j = %d, nums = %d\n", i, j, nums)
			}
			nums[j] = tmp
			fmt.Printf("i = %d, j = %d, nums = %d\n", i, j, nums)
			fmt.Println("---")
		}
	}

	return nums
}

func main() {
	nums := []int{4, 3, 1, 5, 7, 6, 2}
	fmt.Println(InsertionSort(nums))
}

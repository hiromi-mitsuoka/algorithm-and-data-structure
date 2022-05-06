package main

import "fmt"

func ShellSort(nums []int) []int {
	n := len(nums)
	gap := n >> 1

	for gap > 0 {
		for i := gap; i < n; i++ {
			for j := i; j >= gap; j -= gap {
				if nums[j-gap] > nums[j] {
					nums[j], nums[j-gap] = nums[j-gap], nums[j]
				}
			}
		}
		gap >>= 1
	}
	return nums
}

func main() {
	nums := []int{4, 3, 1, 5, 7, 6, 2}
	fmt.Println(ShellSort(nums))
}

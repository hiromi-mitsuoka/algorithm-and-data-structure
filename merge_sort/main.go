package main

import "fmt"

var center int
var left, right []int

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left := MergeSort(nums[:len(nums)/2])
	right := MergeSort(nums[len(nums)/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	ans := []int{}
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		ans = append(ans, left[i])
	}

	for ; j < len(right); j++ {
		ans = append(ans, right[j])
	}

	return ans
}

func main() {
	nums := []int{2, 4, 3, 1, 5, 7, 6}
	fmt.Println(MergeSort(nums))
}

// https://blog.boot.dev/golang/merge-sort-golang/

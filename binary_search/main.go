package main

import (
	"fmt"
	"sort"
)

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

	// Use package
	// https://pkg.go.dev/sort#SearchInts
	x := 4
	i := sort.SearchInts(nums, x)
	fmt.Printf("found %d at index %d %v\n", x, i, nums)

	// Check
	y := 7
	j := CheckSearchInts(nums, y)
	fmt.Printf("found %d at index %d %v\n", y, j, nums)
}

// ------------------------------------------------------
// Check methods (SearchInts, Search)

// https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/sort/search.go;l=83
// SearchInts searches for x in a sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
//
func CheckSearchInts(a []int, x int) int {
	return CheckSearch(len(a), func(i int) bool { return a[i] >= x })
}

// https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/sort/search.go;drc=refs%2Ftags%2Fgo1.18.1;l=58
// Search uses binary search
func CheckSearch(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i <= h < j
		if !f(h) { // func(i int) bool { return a[i] >= x }
			i = h + 1 // preserves f(i-1) == fasle
			fmt.Printf("preserves f(i-1) == fasle: !f(h) = %t, i = %d\n", !f(h), i)
		} else {
			j = h // preserves f(j) == true
			fmt.Printf("preserves f(j) == true: !f(h) = %t, j = %d\n", !f(h), j)
		}
	}
	// i == j, f(i-1) == false, and f(j)(=f(i)) == true => answer is i.
	return i
}

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

	// slice := []int{2, 4, 3, 1, 5, 7, 6}
	// fmt.Println(slice)
	// fmt.Println(sort.IntSlice(slice))
}

// // ------------------------------------------------------
// // Check sort package

// func CheckInts(x []int) { CheckSort(CheckIntSlice(x))}

// type CheckIntSlice []int

// func (x CheckIntSlice) CheckSort() { CheckSort(x) }
// func (x CheckIntSlice) Len() int { return(x) }

// func CheckSort(data interface) {
// 	n := data.Len()
// 	CheckQuickSort(data, 0, n, CheckMaxDepth(n))
// }

// func CheckMaxDepth(){
// 	var depth int
// 	for i := n; i > 0; i >>= 1{
// 		depth++
// 	}
// 	return depth*2
// }

// func CheckQuickSort(data interface, a, b, maxDepth int){
// 	// Use ShellSort for slices <= 12 elements
// 	for b - a > 12 {
// 		if maxDepth == 0 {
// 			heapSort(data, a, b)
// 			return
// 		}
// 		maxDepth--
// 		mlo, mhi := CheckDoPivot(data, a, b)
// 		// Avoiding recursion on the larger subproblem guarantees
// 		// a stack depth of at most lg(b-a)
// 		if mlo-a < b-mhi {
// 			CheckQuickSort(data, a, mlo, maxDepth)
// 			a = mhi // i.e., CheckQuickSort(data, mhi, b)
// 		} else {
// 			CheckQuickSort(data, mhi, b, maxDepth)
// 			b = mlo // i.e., CheckQuickSort(data, a, mlo)
// 		}
// 	}

// 	// Do ShellSort pass with gap 6
// 	// It could be written in this simplidied form cause b -a <= 12
// 	for b - a > 1 {
// 		for i:= a+6; i< b; i++ {
// 			if data.Less(i, i-6) {
// 				data.Swap(i, i-6)
// 			}
// 		}
// 		insertionSort(data)
// 	}
// }

// func CheckDoPivot(data interface, lo, hi int) (midlo, midhi int) {

// }

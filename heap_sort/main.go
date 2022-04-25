package main

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {

}

// // Get the parent index
// func parent(i int) int {
// 	return (i - 1) / 2
// }

// // Get the left child index
// func left(i int) int {
// 	return 2*i + 1
// }

// // Get the right child index
// func right(i int) int {
// 	return 2*i + 2
// }

// // Swap keys in the array
// func (h *MaxHeap) swap(i1, i2 int) {
// 	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
// }

func main() {
	// nums := []int{2, 4, 3, 1, 5, 7, 6}
	// fmt.Println(HeapSort(nums))
}

// https://www.youtube.com/watch?v=3DYIgTC4T1o

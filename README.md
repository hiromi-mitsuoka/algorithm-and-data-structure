# algorithm-and-data-structure

## Sort Complexity

| Sort | Complexity |
|------|------------|
| bubble | O(n**2) |
| selection | O(n**2) |
| shell | Interval-dependent |
| insertion | O(n**2) |
| quick | O(n*logn) |
| merge | O(n*logn) |
| heap | O(n*logn) |


## Search Complexity

| Search | Complexity | note |
|--------|------------| ---- |
| linear | O(nm) | Search m data from n data |
| binary | O(log2(n)) |  |

## Sort Feature

### Bubble sort

- Sort to align adjacent elements by comparing their sizes.
- While the algorithm is simple and easy to implement, the worst-case time complexity is O(n**2) and slow.

( https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%96%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88 )

### Selection sort

- Sort by finding the smallest value in the array and replacing it with the first element in the array.
- Sine the worst-case time complexity is O(n**2), a faster method such as quick sort is generally used.
  - However, when other fast methods cannot be used due to limited space, or when the array to sort is smallenough and selective sorting is guaranteed to be fast, it is used.

( https://ja.wikipedia.org/wiki/%E9%81%B8%E6%8A%9E%E3%82%BD%E3%83%BC%E3%83%88#:~:text=%E9%81%B8%E6%8A%9E%E3%82%BD%E3%83%BC%E3%83%88%EF%BC%88%E8%8B%B1%3A%20selection%20sort,%E6%96%B9%E6%B3%95%E3%81%8C%E5%88%A9%E7%94%A8%E3%81%95%E3%82%8C%E3%82%8B%E3%80%82 )


### Shell sort

- It is a generalization of insertion sort.
- It is an algorithm that performs insertion sort for each pair of elements that are some distance apart in the array, and repeats the same sort while decreasing the interval to speed up the prosess.

( https://ja.wikipedia.org/wiki/%E3%82%B7%E3%82%A7%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88 )


### Insertion sort

- Slow compared to quicksorts, merge sorts, etc. but
  - Simple algorithm and easy to implement
  - Fast for small arrays

( https://ja.wikipedia.org/wiki/%E6%8C%BF%E5%85%A5%E3%82%BD%E3%83%BC%E3%83%88 )


### Quick sort

- Sorting algorithm developed by Antony Hoare in 1960
- A type of devide-and-conquer low（分割統治法）.
- It is generary said to be the fastest compared to other sorting methods,
  - but it is not always fast depending on the sequence of target data or the number of data, and the worst-case computational complexity is O(n**2)

( https://ja.wikipedia.org/wiki/%E3%82%AF%E3%82%A4%E3%83%83%E3%82%AF%E3%82%BD%E3%83%BC%E3%83%88 )


### Merge sort

- A bottom-up divide-and-conquer method （分割統治法） where multiple columns that are already aligned are merged into a single column, and the new column is also aligned if the smaller columns are placed in the new column first.
- Compared to quick sort, the worst-case computational complexity is less. Quick sort is usually faster for random

( https://ja.wikipedia.org/wiki/%E3%83%9E%E3%83%BC%E3%82%B8%E3%82%BD%E3%83%BC%E3%83%88 )


### Heap sort

- Takes elements from the unaligned list and adds them to the heap in order. Repeat until all elements are added.
- Extract the root (largest or smallest value) and add it to the aligned list. Iterate until all elements are retrieved.

( https://ja.wikipedia.org/wiki/%E3%83%92%E3%83%BC%E3%83%97%E3%82%BD%E3%83%BC%E3%83%88 )


## Search Feature

### Linear search （線形探索）

- When searching for data in a list or array, the comparison is done from the top, and if it is found, it is done.

( https://ja.wikipedia.org/wiki/%E7%B7%9A%E5%9E%8B%E6%8E%A2%E7%B4%A2#:~:text=%E7%B7%9A%E5%9E%8B%E6%8E%A2%E7%B4%A2%EF%BC%88%E3%81%9B%E3%82%93%E3%81%91%E3%81%84%E3%81%9F%E3%82%93,%E3%81%8C%E8%A6%8B%E3%81%A4%E3%81%8B%E3%82%8C%E3%81%B0%E7%B5%82%E4%BA%86%E3%81%99%E3%82%8B%E3%80%82 )


### Binary search （二分探索）

- Searching for sorted arrays.
- Look at the median value and use the relationship between the median value and the value you want to search for to determine if the value you want to search for is to the right or left of the median value, making sure it doesn't exist on one side.

( https://ja.wikipedia.org/wiki/%E4%BA%8C%E5%88%86%E6%8E%A2%E7%B4%A2#:~:text=%E4%BA%8C%E5%88%86%E6%8E%A2%E7%B4%A2%EF%BC%88%E3%81%AB%E3%81%B6%E3%82%93%E3%81%9F%E3%82%93,%E6%8E%A2%E7%B4%A2%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%AE%E4%B8%80%E3%81%A4%E3%80%82 )


## Linked list Feature

### Singly linked list （片方向リスト）

- This link points to the next node on the list, and if it is the tail of the list, it stores a null value or points to an empty list.


( https://ja.wikipedia.org/wiki/%E9%80%A3%E7%B5%90%E3%83%AA%E3%82%B9%E3%83%88#%E7%89%87%E6%96%B9%E5%90%91%E3%83%AA%E3%82%B9%E3%83%88 )

### Doubly linked list （双方向リスト）

- Each node has two links, one pointing to the next node and the other pointing to the node behind.

( https://ja.wikipedia.org/wiki/%E9%80%A3%E7%B5%90%E3%83%AA%E3%82%B9%E3%83%88#%E5%8F%8C%E6%96%B9%E5%90%91%E3%83%AA%E3%82%B9%E3%83%88 )


## Stack and Queue Features

### Stack

- LIFO (Last in first out)

- In paricular, as a concreate example, it is extremely useful for supporting interrupts and subroutines.

( https://ja.wikipedia.org/wiki/%E3%82%B9%E3%82%BF%E3%83%83%E3%82%AF )

### Queue

- FIFO (First in first out)

- used for handling events or messages in windowing systems, managing processes, and other cases where data needs to be processed in the order in thich it is enterd.
- When th execution time of individual tasks is unpredictable or takes too long to execute imediately, you can use a queue to store tasks and retrieve them later for asynchronous execution.

( https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%A5%E3%83%BC_(%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF) )


## Graph

- An abstract data type consisting of a group of nodes (vertices) and a group of edges (branches) that represent the connection relationship between nodes.
- There are two main types of data structures for the actual representation of graphs.
  - adjacent list
    - It is suitable for sparse graphs
  - adjacency matrix
- Typical operations on graphs
  - Depth-first search（深さ優先探索）
  - Width-first search（幅優先探索）
  - Dijkstra method（ダイクストラ法）
  - Walsall-Floyd method（ワーシャル-フロイド法）

( https://ja.wikipedia.org/wiki/%E3%82%B0%E3%83%A9%E3%83%95_(%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0) )


## Dynamic Programming (DP: 動的計画法) Feature

- This is a general term for a method of dividing a target problem into multiple subproblems and solving them while recording the calculation results of the subproblems.
- It is a generic term for algorithms that satisfy the following two conditions
  1. Use of inductive relationships
    - use the solution of a smaller problem example or the results of a calculation to solve a larger problem example using inductive relationships.
  2. Recording the results of calculations
    - Record the results of calculations starting from small problem examples and calculation results to avoid performing the same calculation many times.
    - For efficient reference in inductive relationships, calculation results are kept under headings such as integers, letters and their combinations.

( https://ja.wikipedia.org/wiki/%E5%8B%95%E7%9A%84%E8%A8%88%E7%94%BB%E6%B3%95#:~:text=%E5%8B%95%E7%9A%84%E8%A8%88%E7%94%BB%E6%B3%95%EF%BC%88%E3%81%A9%E3%81%86,%E7%B7%8F%E7%A7%B0%E3%81%97%E3%81%A6%E3%81%93%E3%81%86%E5%91%BC%E3%81%B6%E3%80%82 )


## How to 〇〇

### create a slice

```golang
n := make(type, length, capacity)
```


example

```golang
n := make([]int, 5, 5)
fmt.Println(n)
fmt.Printf("len: %d, cap: %d", len(n), cap(n))

=> [0 0 0 0 0]
=> len: 5, cap: 5
```

### Blank-separated standard input

```golang
var n int
fmt.Scan(&n)
s := make([]int, n)
for i := range s {
  fmt.Scan(&s[i])
}
fmt.Println(n, s)
```

input
```
5
1 2 3 4 5
```

output
```
5 [1 2 3 4 5]
```


### Create 2D slice with the number of elements specified by variables

( https://qiita.com/ta7uw/items/387f7b81bb7798915a48 )

```golang
var n int
fmt.Scan(&n)

matrix := make([][]int, n)
for i := range matrix {
  matrix[i] = make([]int, n)
}
```

```txt
$ go run .
4
[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
```

### Split standard input into variable length
- Use scanner (bufio)

```golang
scanner := bufio.NewScanner(os.Stdin)

for i := 0; i < 2; i++ {
  if scanner.Scan() {
    str := scanner.Text()
    arr := strings.Split(str, " ")
    fmt.Println(str, arr)
  }
}
```

```txt
$ go run .
1 2 3
1 2 3 [1 2 3]

2 4 6
2 4 6 [2 4 6]
```


### Use shift of bit

( https://www.flyenginer.com/low/go/go%E3%81%AE%E3%83%93%E3%83%83%E3%83%88%E6%BC%94%E7%AE%97%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6.html#toc5 )

```golang
11 << 1
// 001011 = 11
// 010110 = 22

11 << 2
// 001011 = 11
// 101100 = 44
```

- Shifts each bit to the left by the specified number. The rightmout bit vacated by the shift is set to 0.
  - One bit of left shift doubles the value, two bits quadruples it.

---

```golang
11 >> 1
// 1011 = 11
// 0101 = 5

11 >> 2
// 1011 = 11
// 0010 = 2
```

- Shifts each bit to the right by the specified number. THe leftmost bit vacated by the shift is set to 0.
  - One bit of right shift multiplies the value by a factor of 1/2, two bits by a factor of 1/4.
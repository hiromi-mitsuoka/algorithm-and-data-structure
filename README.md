# algorithm-and-data-structure

## Sort Complexity

| Sort | Complexity |
|------|------------|
| bubble | O(n**2) |
| selection | O(n**2) |
| insertion | O(n**2) |
| quick | O(n*logn) |
| merge | O(n*logn) |


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


## Search Feature

### Linear search （線形探索）

- When searching for data in a list or array, the comparison is done from the top, and if it is found, it is done.

( https://ja.wikipedia.org/wiki/%E7%B7%9A%E5%9E%8B%E6%8E%A2%E7%B4%A2#:~:text=%E7%B7%9A%E5%9E%8B%E6%8E%A2%E7%B4%A2%EF%BC%88%E3%81%9B%E3%82%93%E3%81%91%E3%81%84%E3%81%9F%E3%82%93,%E3%81%8C%E8%A6%8B%E3%81%A4%E3%81%8B%E3%82%8C%E3%81%B0%E7%B5%82%E4%BA%86%E3%81%99%E3%82%8B%E3%80%82 )


### Binary search

- Searching for sorted arrays.
- Look at the median value and use the relationship between the median value and the value you want to search for to determine if the value you want to search for is to the right or left of the median value, making sure it doesn't exist on one side.

( https://ja.wikipedia.org/wiki/%E4%BA%8C%E5%88%86%E6%8E%A2%E7%B4%A2#:~:text=%E4%BA%8C%E5%88%86%E6%8E%A2%E7%B4%A2%EF%BC%88%E3%81%AB%E3%81%B6%E3%82%93%E3%81%9F%E3%82%93,%E6%8E%A2%E7%B4%A2%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%AE%E4%B8%80%E3%81%A4%E3%80%82 )


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
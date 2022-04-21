# algorithm-and-data-structure

## Sort Complexity

| Sort | Complexity |
|------|------------|
| bubble | O(n**2) |
| selection | O(n**2) |


## Search Complexity

| Search | Complexity | note |
|--------|------------| ---- |
| linear | O(nm) | Search m data from n data |

## Feature

### Bubble sort

- Sort to align adjacent elements by comparing their sizes.
- While the algorithm is simple and easy to implement, the worst-case time complexity is O(n**2) and slow.

( https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%96%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88 )

### Selection Sort

- Sort by finding the smallest value in the array and replacing it with the first element in the array.
- Sine the worst-case time complexity is O(n**2), a faster method such as quick sort is generally used.
  - However, when other fast methods cannot be used due to limited space, or when the array to sort is smallenough and selective sorting is guaranteed to be fast, it is used.

( https://ja.wikipedia.org/wiki/%E9%81%B8%E6%8A%9E%E3%82%BD%E3%83%BC%E3%83%88#:~:text=%E9%81%B8%E6%8A%9E%E3%82%BD%E3%83%BC%E3%83%88%EF%BC%88%E8%8B%B1%3A%20selection%20sort,%E6%96%B9%E6%B3%95%E3%81%8C%E5%88%A9%E7%94%A8%E3%81%95%E3%82%8C%E3%82%8B%E3%80%82 )

### Linear search


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

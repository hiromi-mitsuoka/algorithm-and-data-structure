# algorithm-and-data-structure

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


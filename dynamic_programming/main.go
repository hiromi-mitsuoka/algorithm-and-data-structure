package main

import "fmt"

// // ======================================================
// // memo_fibonacci_01

// var memo [50]int

// func memo_fibonacci(n int) int {
// 	if n == 0 || n == 1 {
// 		memo[n] = 1
// 		return memo[n]
// 	}

// 	if memo[n] != -1 {
// 		return memo[n]
// 	}

// 	memo[n] = memo_fibonacci(n-1) + memo_fibonacci(n-2)
// 	return memo[n]
// }

// func main() {
// 	var n int
// 	for i := 0; i < 50; i++ {
// 		memo[i] = -1
// 	}
// 	// fmt.Println(memo)
// 	// fmt.Println(len(memo))

// 	fmt.Scan(&n)
// 	fmt.Printf("memo_fibonacci(%d) = %d\n", n, memo_fibonacci(n))
// 	fmt.Println(memo)
// }

// // https://ja.wikipedia.org/wiki/%E3%83%95%E3%82%A3%E3%83%9C%E3%83%8A%E3%83%83%E3%83%81%E6%95%B0

// ======================================================
// memo_fibonacci_02
// https://www.slideshare.net/iwiwi/ss-3578511

// // Use memo
// const MAX_N int = 50
// var done [MAX_N + 1]bool
// var memo [MAX_N + 1]int

// func memo_fib(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}

// 	if done[n] {
// 		return memo[n]
// 	}

// 	memo[n] = memo_fib(n-1) + memo_fib(n-2)
// 	done[n] = true
// 	return memo[n]
// }

// // Use loop
// const MAX_N int = 50

// var dp [MAX_N + 1]int

// func loop_fib(n int) int {
// 	dp[0], dp[1] = 1, 1

// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}

// 	return dp[n]
// }

// func main() {
// 	// // Use memo
// 	// var n int
// 	// fmt.Scan(&n)

// 	// fmt.Println(memo_fib(n))
// 	// fmt.Println(memo, done)

// 	// Use loop
// 	var n int
// 	fmt.Scan(&n)

// 	fmt.Println(loop_fib(n))
// 	fmt.Println(dp)
// }

// ======================================================
// ナップサック問題
// 重さと価値がそれぞれw_i, v_iであるようなn個の品物があります。
// これらの品物から、重さの総和がWを超えないように選んだときの、
// 価値の総和の最大値を求めなさい。

// https://blog.shirakiyo.dev/post/dynamic-programming/

// // ======================================================
// // Gloval search
// var (
// 	n, W int
// 	w, v []int
// )

// func rec(i, j int) (res int) {
// 	if i == n {
// 		// There are no more goods left
// 		res = 0
// 		////  fmt.Printf("res = 0, i = %d\n", i)
// 		////  fmt.Printf("=============== i == n: res = %d, i = %d, j = %d\n", res, i, j)
// 	} else if j < w[i] {
// 		// This item will not fit (Over)
// 		res = rec(i+1, j)
// 		////  fmt.Printf("Over, res = %d, rec(i+1, j) = rec(%d, %d)\n", res, i+1, j)
// 		////  fmt.Printf("Over: res = %d, i = %d, j = %d ===============\n", res, i, j)
// 	} else {
// 		if rec(i+1, j) < rec(i+1, j-w[i])+v[i] { // Try both without and with item and return the larger one
// 			res = rec(i+1, j-w[i]) + v[i] // Subtract w[i] from the total amount of the second artument.
// 			////  fmt.Printf("Add, i = %d, j = %d, v[i] = %d, w[i] = %d, rec(i+1, j-w[i]) + v[i] = rec(%d, %d) + %d, res = %d\n", i, j, v[i], w[i], i+1, j-w[i], v[i], res)
// 			////  fmt.Printf("res = %d, i = %d, j = %d\n", res, i, j)
// 		} else {
// 			res = rec(i+1, j)
// 			//// fmt.Printf("Add, i = %d, j = %d, v[i] = %d, w[i] = %d, rec(i+1, j) = rec(%d, %d), res = %d\n", i, j, v[i], w[i], i+1, j, res)
// 			////  fmt.Printf("res = %d, i = %d, j = %d\n", res, i, j)
// 		}
// 	}
// 	return
// }

// func main() {
// 	n = 4
// 	w = []int{2, 1, 3, 2}
// 	v = []int{3, 2, 4, 2}

// 	// n = 2
// 	// w = []int{2, 1}
// 	// v = []int{3, 2}
// 	// // (w, v) = {2, 3}, {1, 2}
// 	W = 5
// 	fmt.Println(n, w, v, W)

// 	fmt.Println(rec(0, W))
// }

// ======================================================
// Memorizing the search
var (
	n, W int
	w, v []int
	memo [][]int // memo
)

func rec(i, j int) (res int) {
	if memo[i][j] >= 0 {
		// Use memo
		return memo[i][j]
	}

	if i == n {
		// There are no more goods left
		res = 0
	} else if j < w[i] {
		// This item will not fit (Over)
		res = rec(i+1, j)
	} else {
		if rec(i+1, j) < rec(i+1, j-w[i])+v[i] { // Try both without and with item and return the larger one
			res = rec(i+1, j-w[i]) + v[i] // Subtract w[i] from the total amount of the second artument.
		} else {
			res = rec(i+1, j)
		}
		memo[i][j] = res
	}
	return
}

func main() {
	n = 4
	w = []int{2, 1, 3, 2}
	v = []int{3, 2, 4, 2}
	// n = 2
	// w = []int{2, 1}
	// v = []int{3, 2}
	// // (w, v) = {2, 3}, {1, 2}
	W = 5
	fmt.Println(n, w, v, W)

	memo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j <= W; j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	fmt.Println(rec(0, W))
	fmt.Println(memo)
}

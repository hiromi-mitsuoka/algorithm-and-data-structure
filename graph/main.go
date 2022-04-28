package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// // ============================================
// // Use an adjacency list

// // An adjacency list graph
// type Graph struct {
// 	vertices []*Vertex
// }

// type Vertex struct {
// 	key      int
// 	adjacent []*Vertex
// }

// func (g *Graph) AddVertex(k int) {
// 	if contains(g.vertices, k) {
// 		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
// 		fmt.Println(err.Error())
// 	} else {
// 		g.vertices = append(g.vertices, &Vertex{key: k})
// 	}
// }

// func (g *Graph) AddEdge(from, to int) {
// 	fromVertex := g.vertices[from]
// 	toVertex := g.vertices[to]

// 	if fromVertex == nil || toVertex == nil {
// 		err := fmt.Errorf("Invalid edge (%v ---> %v)", from, to)
// 		fmt.Println(err.Error())
// 	} else if contains(fromVertex.adjacent, to) {
// 		err := fmt.Errorf("Existing edge (%v ---> %v)", from, to)
// 		fmt.Println(err.Error())
// 	} else {
// 		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
// 	}
// }

// func contains(s []*Vertex, k int) bool {
// 	for _, v := range s {
// 		if k == v.key {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (g *Graph) Print() {
// 	for _, v := range g.vertices {
// 		fmt.Printf("\nVertex %v : ", v.key)
// 		for _, v := range v.adjacent {
// 			fmt.Printf("%v", v.key)
// 		}
// 	}
// 	fmt.Println()
// }

// func main() {
// 	test := &Graph{}

// 	for i := 0; i < 5; i++ {
// 		test.AddVertex(i)
// 	}

// 	// test.AddVertex(0)
// 	test.AddEdge(1, 2)
// 	test.AddEdge(1, 2)
// 	test.AddEdge(3, 2)
// 	test.Print()
// }

// // https://www.youtube.com/watch?v=bSZ57h7GN2w

// ============================================
// Use an adjacency matrix

// Problem
// 入力
// 最初の行にG の頂点数n が与えられます。続くn 行で各頂点u の隣接リストA d j [u]が以下の形式で与えられます:
// u k v1 v2 ... vk
// uは頂点の番号、kはuの出次数、v1 v2 ... vk はuに隣接する頂点の番号を示します。

// 出力
// 出力例に従い、G の隣接行列を出力してください。a i j の間に1つの空白を入れてく
// ださい

// 制約 1 ≤ n ≤ 100

func main() {
	var n int
	fmt.Scan(&n)

	// Create 2D array(slice) with the number of elements specified by variables
	// https://qiita.com/ta7uw/items/387f7b81bb7798915a48
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			str := scanner.Text()
			arr := strings.Split(str, " ")
			u, _ := strconv.Atoi(arr[0])
			k, _ := strconv.Atoi(arr[1])
			v_arr := arr[2:]
			// fmt.Println(u, k, v_arr)

			for j := 0; j < k; j++ {
				for idx := range v_arr {
					v, _ := strconv.Atoi(v_arr[idx])
					matrix[u-1][v-1] = 1
				}
			}
		}
	}

	fmt.Println(matrix)
}

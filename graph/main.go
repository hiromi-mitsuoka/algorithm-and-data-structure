package main

import "fmt"

// ============================================
// Use an adjacency list

// An adjacency list graph
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.vertices[from]
	toVertex := g.vertices[to]

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v ---> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v ---> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	// test.AddVertex(0)
	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(3, 2)
	test.Print()
}

// https://www.youtube.com/watch?v=bSZ57h7GN2w

// ============================================
// Use an adjacency matrix

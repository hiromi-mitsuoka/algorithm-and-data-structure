package main

import (
	"fmt"
)

const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	// hashtable := Init()
	// fmt.Println(hashtable)
	// // fmt.Println(hash("TEST"))     // 5
	// // fmt.Println(hash("TESTTEST")) // 3

	// testBucket := &bucket{}
	// testBucket.insert("TEST")
	// testBucket.insert("TESTTEST")

	// fmt.Println("TEST", testBucket.search("TEST"))
	// fmt.Println("TESTTEST", testBucket.search("TESTTEST"))

	// testBucket.delete("TESTTEST")
	// fmt.Println("TESTTEST", testBucket.search("TESTTEST"))

	hashtable := Init()
	list := []string{
		"TEST01",
		"TEST02",
		"TEST03",
		"TEST04",
		"TEST05",
		"TEST06",
		"TEST07",
		"TEST08",
	}

	for _, v := range list {
		hashtable.Insert(v)
	}

	hashtable.Delete("TEST08")

	fmt.Println("TEST08", hashtable.Search("TEST08"))
	fmt.Println("TEST01", hashtable.Search("TEST01"))
}

// https://www.youtube.com/watch?v=zLnJcAt1aKs

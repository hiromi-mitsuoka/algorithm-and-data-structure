package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	currentNode := l.head
	for i := 0; i < l.len; i++ {
		if currentNode.next == nil {
			currentNode.next = &n
			l.len++
			return
		}
		currentNode = currentNode.next
	}
}

// Inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value int) {
	newNode := Node{}
	newNode.value = value

	if pos < 0 {
		return
	}

	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}

	if pos > l.len {
		return
	}

	n := l.GetAt(pos)
	newNode.next = n
	preNode := l.GetAt(pos - 1)
	preNode.next = &newNode
	l.len++
}

func (l *LinkedList) GetAt(pos int) *Node {
	currentNode := l.head

	if pos < 0 {
		return currentNode
	}

	if pos > (l.len - 1) {
		return nil
	}

	for i := 0; i < pos; i++ {
		currentNode = currentNode.next
	}

	return currentNode
}

// Traverse
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	currentNode := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", *currentNode)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) Search(val int) int {
	currentNode := l.head
	for i := 0; i < l.len; i++ {
		if currentNode.value == val {
			return i
		}
		currentNode = currentNode.next
	}
	return -1
}

func (l *LinkedList) DeleteAt(pos int) error {
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}

	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}

	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
}

func (l *LinkedList) DeleteVal(val int) error {
	currentNode := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.len; i++ {
		if currentNode.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = currentNode.next
			}
			l.len--
			return nil
		}
		currentNode = currentNode.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

func main() {
	linkedlist := &LinkedList{}
	fmt.Printf("Initialize linkedlist: %v\n", linkedlist)

	val := 4
	fmt.Printf("===== Insert %d =====\n", val)
	linkedlist.Insert(val)
	fmt.Printf("Insert %d, linkedlist: %v, head: %v, len: %v\n", val, *linkedlist, *linkedlist.head, *&linkedlist.len)

	val = 8
	fmt.Printf("===== Insert %d =====\n", val)
	linkedlist.Insert(val)
	fmt.Printf("Insert %d, linkedlist: %v, head: %v, len: %v\n", val, *linkedlist, *linkedlist.head, *&linkedlist.len)

	pos := 1
	fmt.Printf("===== GetAt %d =====\n", pos)
	fmt.Printf("val is %v\n", *&linkedlist.GetAt(pos).value)

	val = 9
	pos = 1
	fmt.Printf("===== InsertAt %d, val = %d =====\n", pos, val)
	linkedlist.InsertAt(pos, val)
	linkedlist.Print()

	val = 9
	fmt.Printf("===== Search %d, position is %d =====\n", val, linkedlist.Search(val))
	linkedlist.Print()

	pos = 2
	fmt.Printf("===== DeleteAt %d =====\n", pos)
	linkedlist.DeleteAt(pos)
	linkedlist.Print()

	val = 9
	fmt.Printf("===== DeleteVal %d =====\n", val)
	linkedlist.DeleteVal(val)
	linkedlist.Print()
}

// https://levelup.gitconnected.com/go-singly-linked-lists-with-insertion-deletion-traversal-e9da5bb0dbe1

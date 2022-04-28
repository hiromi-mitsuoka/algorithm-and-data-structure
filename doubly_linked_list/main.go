package main

import "fmt"

type Node struct {
	data       int
	prev, next *Node
}

type DoublyLinkedList struct {
	len        int
	head, tail *Node
}

func initDoublyList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) AddFrontNode(data int) {
	newNode := &Node{
		data: data,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}

	d.len++
	return
}

func (d *DoublyLinkedList) AddEndNode(data int) {
	newNode := &Node{
		data: data,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}

	d.len++
	return
}

func (d *DoublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}

	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.next
	}

	fmt.Println()
	return nil
}

func (d *DoublyLinkedList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}

	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}

	fmt.Println()
	return nil
}

func (d *DoublyLinkedList) Size() int {
	return d.len
}

func main() {
	doublyList := initDoublyList()

	// fmt.Println("AddFrontNode")
	doublyList.AddFrontNode(3)
	doublyList.AddFrontNode(6)

	// fmt.Println("AddEndNode")
	doublyList.AddEndNode(8)
	doublyList.AddEndNode(10)
	doublyList.AddEndNode(12)
	doublyList.AddEndNode(14)

	fmt.Printf("Size of doubly linked list: %d\n", doublyList.Size())

	err := doublyList.TraverseForward()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = doublyList.TraverseReverse()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// https://golangbyexample.com/doubly-linked-list-golang/

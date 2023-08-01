package main

import "fmt"

// Node
type Node struct {
	data     int   // value of node
	nextNode *Node // reference of next node
}

// Singly linked list
type LinkedList struct {
	head *Node
}

// append
func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.nextNode != nil {
			current = current.nextNode
		}
		current.nextNode = newNode
	}
}

// Print
func (list *LinkedList) Print() {
	current := list.head

	for current != nil {
		fmt.Printf("%d", current.data)
		current = current.nextNode
	}

	fmt.Println()
}

func main() {
	list := &LinkedList{}
	listTwo := &LinkedList{}

	list.Append(2)
	list.Append(4)
	list.Append(3)

	listTwo.Append(5)
	listTwo.Append(6)
	listTwo.Append(4)

	list.Print()
	listTwo.Print()
}

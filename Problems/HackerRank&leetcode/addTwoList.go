/* Program to merge two list */

package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	current := list1 	

	for current.Next != nil {
		current = current.Next
	}

	current.Next = list2

	return list1
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

	fmt.Println("")
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	// list1 := &ListNode{}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	// list2 := &ListNode{}

	mergredList := mergeTwoLists(list1, list2)
	printList(mergredList)
}

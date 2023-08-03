package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeAndSortList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	current :=  list1

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
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

	mergredList := mergeAndSortList(list1, list2)
	printList(mergredList)
}

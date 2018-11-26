package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
 	Val int
 	Next *ListNode
 }

func deleteDuplicates(head *ListNode) *ListNode {

	var prevNode *ListNode
	var curNode *ListNode

	curNode = head

	for curNode != nil {
		if (prevNode != nil && prevNode.Val == curNode.Val) {
			prevNode.Next = curNode.Next
		} else {
			prevNode = curNode
		}
		curNode = curNode.Next
	}
	return head
}

func main() {
	l14 := &ListNode{
		Val: 4,
		Next: nil,
	}

	l13 := &ListNode{
		Val: 3,
		Next: l14,
	}

	l12 := &ListNode{
		Val: 3,
		Next: l13,
	}
	l11 := &ListNode{
		Val: 1,
		Next: l12,
	}
	l1 := &ListNode{
		Val: 1,
		Next: l11,
	}

	newList := deleteDuplicates(l1)

	for newList != nil {
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}


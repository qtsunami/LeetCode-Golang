package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	nextNode := head

	for nextNode != nil && head != nil && nextNode.Next != nil {
		head = head.Next
		nextNode = nextNode.Next.Next

		if head == nextNode {
			return nextNode
		}
	}

	return nil
}

func main() {
	l3 := &ListNode{
		Val: 3,
		Next: nil,
	}

	l2 := &ListNode{
		Val: 2,
		Next: l3,
	}

	l1 := &ListNode{
		Val: 1,
		Next: l2,
	}

	l0 := &ListNode{
		Val: 0,
		Next: l1,
	}

	l3.Next = l2

	//fmt.Println(l0)
	//
	//l4 := &ListNode{
	//	Val: 0,
	//	Next:nil,
	//}
	fmt.Println(detectCycle(l0))
}



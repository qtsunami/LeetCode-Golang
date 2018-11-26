package main

import "fmt"

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？


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

func reverseList(head *ListNode) *ListNode {

	var prevNode *ListNode
	curNode := head
	for curNode != nil {
		prevNode, curNode.Next, curNode = curNode, prevNode, curNode.Next
	}

	return prevNode
}

func main() {
	l14 := &ListNode{
		Val: 5,
		Next: nil,
	}

	l13 := &ListNode{
		Val: 4,
		Next: l14,
	}

	l12 := &ListNode{
		Val: 3,
		Next: l13,
	}
	l11 := &ListNode{
		Val: 2,
		Next: l12,
	}
	l1 := &ListNode{
		Val: 1,
		Next: l11,
	}


	rNode := reverseList(l1)

	for rNode != nil {
		fmt.Println(rNode.Val)
		rNode = rNode.Next
	}
}

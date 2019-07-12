package main

import "fmt"

/**
 * 合并两个有序链表
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 声明一下新的节点
	var newNode *ListNode
	// 声明新节点的前置结点
	var preNode *ListNode
	// 声明一个头结点
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	var val int
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l2.Val
			l2 = l2.Next
		}
		newNode = &ListNode{
			Val:  val,
			Next: nil,
		}
		// 首节点的Next是否为空
		if head.Next == nil {
			preNode = newNode
			head.Next = newNode
		} else {
			preNode.Next = newNode
			preNode = newNode
		}

	}
	// 如果 l1 有剩余则直接接后面即可
	if l1 != nil {
		preNode.Next = l1
	}
	// 如果 l2 有剩余则直接接后面即可
	if l2 != nil {
		preNode.Next = l2
	}
	return head.Next
}

func main() {
	l12 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l11 := &ListNode{
		Val:  2,
		Next: l12,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l11,
	}

	l22 := &ListNode{
		Val:  4,
		Next: nil,
	}

	l21 := &ListNode{
		Val:  3,
		Next: l22,
	}

	l2 := &ListNode{
		Val:  1,
		Next: l21,
	}

	mergeLink := mergeTwoLists(l1, l2)

	//fmt.Println(mergeLink)

	for mergeLink != nil {
		fmt.Println(mergeLink)
		mergeLink = mergeLink.Next
	}
}

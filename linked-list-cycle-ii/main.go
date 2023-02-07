package main

import "fmt"

//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//不允许修改 链表。

type ListNode struct {
	Val  int
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
		Val:  3,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}

	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	l0 := &ListNode{
		Val:  0,
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

package main

import "fmt"

/**
 * 请判断一个链表是否为回文链表。
 * 示例 1: 输入: 1->2   输出: false
 * 示例 2: 输入: 1->2->2->1  输出: true
 * 示例 3: 输入: 1  输出: true
 * 示例 4: 输入: 1->2->1  输出: true
 */

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

func addNode(node *ListNode, num int) {
	//fmt.Println("addNode", num)
	addNode := ListNode{
		Val:  num,
		Next: nil,
	}

	node.Next = &addNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var setSlice []int
	node := head.Next
	setSlice = append(setSlice, head.Val)
	for node != nil {
		setSlice = append(setSlice, node.Val)
		node = node.Next
	}

	len := len(setSlice)

	for i := 0; i < len/2; i++ {
		if setSlice[i] != setSlice[len-i-1] {
			return false
		}
	}

	return true
}

func main() {
	//head := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//
	//addNode(head, 2)
	//addNode(head.Next, 3)
	//addNode(head.Next.Next, 3)
	//addNode(head.Next.Next.Next, 2)
	//addNode(head.Next.Next.Next.Next, 1)

	var head *ListNode
	isPalind := isPalindrome(head)
	fmt.Println(isPalind)

}

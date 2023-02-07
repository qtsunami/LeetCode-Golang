package main

import (
	"fmt"
)

// 节点
type Node struct {
	val  int
	next *Node
}

// 链接表
type MyLinkedList struct {
	head   *Node
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.length-1 {
		return -1
	}

	node := this.head

	for i := 0; i < index && node != nil; i++ {
		node = node.next
	}

	if node == nil {
		return -1
	}

	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.head

	newFirstNode := &Node{
		val:  val,
		next: head,
	}

	this.head = newFirstNode
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}

	curNode := this.head
	for curNode.next != nil {
		curNode = curNode.next
	}

	curNode.next = &Node{
		val:  val,
		next: nil,
	}
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	if index < 0 || index > this.length-1 {
		return
	}

	curNode := this.head
	for i := 0; i < index-1; i++ {
		curNode = curNode.next
	}

	newNode := &Node{val: val, next: curNode.next}
	curNode.next = newNode
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.length-1 {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.length--
	}
	curNode := this.head
	for i := 0; i < index-1; i++ {
		curNode = curNode.next
	}

	curNode.next = curNode.next.next
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	obj := Constructor()
	obj.AddAtHead(11)
	obj.AddAtTail(10)
	fmt.Println(obj.head)
	//fmt.Println(obj.Get(1))
	//obj.AddAtIndex(1, 4)
	//obj.DeleteAtIndex(2);

	//fmt.Println(obj.val)
	//first := obj.next
	//for first != nil {
	//	fmt.Println(first.val)
	//	first = first.next
	//}
}

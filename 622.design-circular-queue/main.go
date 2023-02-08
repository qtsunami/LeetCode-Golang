package main

type Node struct {
	value int
	next  *Node
}

type MyCircularQueue struct {
	headNode *Node
	rearNode *Node
	size     int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		headNode: nil,
		rearNode: nil,
		size:     0,
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {

	if this.IsFull() {
		return false
	}

	enNode := Node{
		value: value,
		next:  nil,
	}

	// If headNode is nil, headNode and rearNode equals enNode
	if this.IsEmpty() {
		this.headNode = &enNode
		this.rearNode = &enNode
	} else {
		// origin rearNode next point is enNode， and new rearNode is EnNode
		this.rearNode.next = &enNode
		this.rearNode = &enNode
	}

	this.size++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.headNode = this.headNode.next
	this.size--

	if this.IsEmpty() {
		this.rearNode = nil
	}
	return true
}

func (this *MyCircularQueue) Front() int {

	if this.IsEmpty() {
		return -1
	}
	return this.headNode.value
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.rearNode.value
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {

}

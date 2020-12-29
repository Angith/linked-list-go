package data

import "fmt"

// Insert elements in linkedlist
func (head *Node) Insert(nv chan int) {
	var nodePtr *Node
	var lastPtr *Node
	for val := range nv {
		nodePtr = NewNode(val)
		if lastPtr != nil {
			lastPtr.Next = nodePtr
		} else {
			head.Next = nodePtr
		}
		lastPtr = nodePtr
		nodePtr = lastPtr.Next
	}
}

// Traverse display elements in linkedlist
func (head *Node) Traverse() {
	node := head
	for {
		fmt.Println(node.Data)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

package cmd

import "linked-list/app/data"

// Linkedlist main app
type Linkedlist struct {
	data.LinkedList
}

// NewLinkedList create a linked list with a head
func NewLinkedList(headVal int) Linkedlist {
	head := data.NewNode(headVal)
	return Linkedlist{
		LinkedList: head,
	}
}

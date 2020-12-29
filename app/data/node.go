package data

// Node defines a node of linked list
type Node struct {
	Data int
	Next *Node
}

// NewNode return instance of Node
func NewNode(val int) *Node {
	return &Node{
		Data: val,
		Next: nil,
	}
}

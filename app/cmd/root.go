package cmd

// Execute starts execution
func Execute() {
	nv := make(chan int)
	go readData(nv)
	headVal := <-nv
	head := NewLinkedList(headVal)
	head.Insert(nv)
	head.Traverse()
}

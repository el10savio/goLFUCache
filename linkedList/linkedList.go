package linkedlist

// Node ...
type Node struct {
	Value    interface{}
	previous *Node
	next     *Node
}

// DoublyLinkedList ...
func DoublyLinkedList() *Node {
	return &Node{}
}

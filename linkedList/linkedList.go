package linkedlist

// Node ...
type Node struct {
	Value    interface{}
	previous *Node
	next     *Node
}

// DoublyLinkedList ...
type DoublyLinkedList struct {
	Length uint
	head   *Node
}

// InitDoublyLinkedList ...
func InitDoublyLinkedList(length uint) *DoublyLinkedList {
	return &DoublyLinkedList{Length: length, head: nil}
}

package linkedlist

import "errors"

// Node ...
type Node struct {
	Value    int
	previous *Node
	next     *Node
}

// DoublyLinkedList ...
type DoublyLinkedList struct {
	Length uint
	head   *Node
	tail   *Node
}

// InitDoublyLinkedList ...
func InitDoublyLinkedList(length uint) *DoublyLinkedList {
	return &DoublyLinkedList{Length: length, head: nil, tail: nil}
}

// List ...
func (dList *DoublyLinkedList) List() []int {
	return traverseList(dList.head)
}

// traverseList ...
func traverseList(head *Node) []int {
	if head == nil {
		return []int{}
	}
	return append([]int{head.Value}, traverseList(head.next)...)
}

// GetHead ...
func (dlist *DoublyLinkedList) GetHead() (int, error) {
	if dlist.head == nil {
		return 0, errors.New("head does not exist")
	}
	return dlist.head.Value, nil
}

// GetTail ...
func (dlist *DoublyLinkedList) GetTail() (int, error) {
	if dlist.tail == nil {
		return 0, errors.New("tail does not exist")
	}
	return dlist.tail.Value, nil
}

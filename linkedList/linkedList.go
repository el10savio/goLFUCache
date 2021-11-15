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

// InsertAtFront ...
func (dList *DoublyLinkedList) InsertAtFront(value int) {
	node := Node{Value: value}

	if dList.head == nil {
		dList.head = &node
		return
	}

	temp := dList.head.next
	dList.head = &node
	temp.previous = dList.head
}

// InsertAtEnd ...
func (dList *DoublyLinkedList) InsertAtEnd(value int) {
	node := Node{Value: value}

	if dList.tail == nil {
		dList.tail = &node
		return
	}

	temp := dList.tail.previous
	dList.tail = &node
	temp.next = dList.tail
}

// InsertBetween ...
func (dList *DoublyLinkedList) InsertBetween(value, valuePrevious, valueNext int) error {
	head := dList.head

	if head == nil {
		return errors.New("head is not present")
	}

	node := Node{Value: value}

	for head != nil && head.next != nil {
		if head.Value == valuePrevious && (head.next).Value == valueNext {
			node.previous, node.next = head, head.next
			(head.next).previous = &node
			head.next = &node
			return nil
		}
		head = head.next
	}

	return errors.New("range not present")
}

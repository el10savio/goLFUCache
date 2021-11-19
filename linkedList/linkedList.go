package linkedlist

import (
	"errors"
)

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
func InitDoublyLinkedList(length uint) (*DoublyLinkedList, error) {
	if length == 0 {
		return nil, errors.New("list length cannot be zero")
	}
	return &DoublyLinkedList{Length: length, head: nil, tail: nil}, nil
}

// List ...
func (dList *DoublyLinkedList) List() []int {
	return traverseList(dList.head)
}

// ListReverse ...
func (dList *DoublyLinkedList) ListReverse() []int {
	return traverseListReverse(dList.tail)
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

// Prepend (InsertAtFront) ...
func (dList *DoublyLinkedList) Prepend(value int) {
	node := Node{Value: value}

	if dList.head == nil {
		dList.head = &node
		dList.tail = &node
		return
	}

	dList.head.previous = &node
	node.next = dList.head
	dList.head = &node
}

// Append (InsertAtEnd) ...
func (dList *DoublyLinkedList) Append(value int) {
	node := Node{Value: value}

	if dList.head == nil {
		dList.head = &node
		dList.tail = &node
		return
	}

	dList.tail.next = &node
	node.previous = dList.tail
	dList.tail = &node
}

// InsertBetween ...
func (dList *DoublyLinkedList) InsertBetween(value, valuePrevious, valueNext int) error {
	head := dList.head

	if head == nil {
		return errors.New("head does not exist")
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

	return errors.New("range does not exist")
}

// RemoveHead ...
func (dList *DoublyLinkedList) RemoveHead() error {
	if dList.head == nil {
		return errors.New("head does not exist")
	}

	if dList.head == dList.tail {
		dList.head = nil
		dList.tail = nil
		return nil
	}

	dList.head = dList.head.next
	(dList.head).previous = nil

	return nil
}

// RemoveTail ...
func (dList *DoublyLinkedList) RemoveTail() error {
	if dList.tail == nil {
		return errors.New("tail does not exist")
	}

	if dList.tail == dList.head {
		dList.head = nil
		dList.tail = nil
		return nil
	}

	dList.tail = dList.tail.previous
	(dList.tail).next = nil

	return nil
}

// RemoveElement ...
func (dList *DoublyLinkedList) RemoveElement(value int) error {
	if dList.head == nil {
		return errors.New("head does not exist")
	}

	node, err := dList.FindElement(value)
	if err != nil {
		return err
	}

	if node == dList.head {
		return dList.RemoveHead()
	}

	if node == dList.tail {
		return dList.RemoveTail()
	}

	node.previous.next = node.next
	node.next.previous = node.previous

	return nil
}

// FindElement ...
func (dList *DoublyLinkedList) FindElement(value int) (*Node, error) {
	if dList.head == nil {
		return nil, errors.New("head does not exist")
	}

	temp := dList.head

	for temp != nil {
		if temp.Value == value {
			return temp, nil
		}
		temp = temp.next
	}

	return nil, errors.New("element does not exist")
}

// ClearList ...
func (dList *DoublyLinkedList) ClearList() {
	dList = &DoublyLinkedList{Length: dList.Length, head: nil, tail: nil}
}

// GetNextElement ...
func GetNextElement(node *Node) (*Node, error) {
	if node == nil {
		return nil, errors.New("element does not exist")
	}

	if node.next == nil {
		return nil, errors.New("next element does not exist")
	}

	return node.next, nil
}

// GetPreviousElement ...
func GetPreviousElement(node *Node) (*Node, error) {
	if node == nil {
		return nil, errors.New("element does not exist")
	}

	if node.previous == nil {
		return nil, errors.New("previous element does not exist")
	}

	return node.previous, nil
}

// traverseList ...
func traverseList(head *Node) []int {
	if head == nil {
		return []int{}
	}
	return append([]int{head.Value}, traverseList(head.next)...)
}

// traverseListReverse ...
func traverseListReverse(tail *Node) []int {
	if tail == nil {
		return []int{}
	}
	return append([]int{tail.Value}, traverseListReverse(tail.previous)...)
}

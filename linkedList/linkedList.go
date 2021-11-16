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

	// if dList.head.next == nil {
	// 	dList.head = dList.head.next
	// 	return nil
	// }

	dList.head = dList.head.next

	return nil
}

// RemoveTail ...
func (dList *DoublyLinkedList) RemoveTail() error {
	if dList.tail == nil {
		return errors.New("tail does not exist")
	}

	// if dList.tail.previous == nil {
	// 	dList.tail = dList.tail.previous
	// 	return nil
	// }

	dList.tail = dList.tail.previous

	return nil
}

// RemoveElement ...
func (dList *DoublyLinkedList) RemoveElement(value int) error {
	head := dList.head

	if head == nil {
		return errors.New("head does not exist")
	}

	for head != nil && head.next != nil {
		if head.next.Value == value {
			head.next = head.next.next
			head.next.previous = head
			return nil
		}
		head = head.next
	}

	return errors.New("node element does not exist")
}

// ClearList ...
func (dList *DoublyLinkedList) ClearList() {
	dList = &DoublyLinkedList{Length: dList.Length, head: nil, tail: nil}
}

// traverseList ...
func traverseList(head *Node) []int {
	if head == nil {
		return []int{}
	}
	return append([]int{head.Value}, traverseList(head.next)...)
}

// GetNextElement ...
func GetNextElement(head *Node) (int, error) {
	if head == nil {
		return 0, errors.New("node element does not exist")
	}

	if head.next == nil {
		return 0, errors.New("node next element does not exist")
	}

	return head.next.Value, nil
}

// GetPreviousElement ...
func GetPreviousElement(head *Node) (int, error) {
	if head == nil {
		return 0, errors.New("node element does not exist")
	}

	if head.previous == nil {
		return 0, errors.New("node previous element does not exist")
	}

	return head.previous.Value, nil
}

package linkedlist

// Node ...
type Node struct {
	Value    int
	previous *Node
	next     *Node
}

// DoublyLinkedList ...
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// InitDoublyLinkedList ...
func InitDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil}
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
func (dList *DoublyLinkedList) GetHead() (int, error) {
	if dList.head == nil {
		return 0, errHeadDoesNotExist
	}
	return dList.head.Value, nil
}

// GetTail ...
func (dList *DoublyLinkedList) GetTail() (int, error) {
	if dList.tail == nil {
		return 0, errTailDoesNotExist
	}
	return dList.tail.Value, nil
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
		return errHeadDoesNotExist
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

	return errRangeDoesNotExist
}

// RemoveHead ...
func (dList *DoublyLinkedList) RemoveHead() error {
	if dList.head == nil {
		return errHeadDoesNotExist
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
		return errTailDoesNotExist
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
		return errHeadDoesNotExist
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
		return nil, errHeadDoesNotExist
	}

	temp := dList.head

	for temp != nil {
		if temp.Value == value {
			return temp, nil
		}
		temp = temp.next
	}

	return nil, errElementDoesNotExist
}

// ClearList ...
func (dList *DoublyLinkedList) ClearList() {
	dList = &DoublyLinkedList{head: nil, tail: nil}
}

// GetNextElement ...
func GetNextElement(node *Node) (*Node, error) {
	if node == nil {
		return nil, errElementDoesNotExist
	}

	if node.next == nil {
		return nil, errNextElementDoesNotExist
	}

	return node.next, nil
}

// GetPreviousElement ...
func GetPreviousElement(node *Node) (*Node, error) {
	if node == nil {
		return nil, errElementDoesNotExist
	}

	if node.previous == nil {
		return nil, errPreviousElementDoesNotExist
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

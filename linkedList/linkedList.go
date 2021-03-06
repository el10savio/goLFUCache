package linkedlist

// Node ...
type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

// DoublyLinkedList ...
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// InitDoublyLinkedList ...
func InitDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{Head: nil, Tail: nil}
}

// List ...
func (dList *DoublyLinkedList) List() []int {
	return traverseList(dList.Head)
}

// ListReverse ...
func (dList *DoublyLinkedList) ListReverse() []int {
	return traverseListReverse(dList.Tail)
}

// GetHead ...
func (dList *DoublyLinkedList) GetHead() (int, error) {
	if dList.Head == nil {
		return 0, errHeadDoesNotExist
	}
	return dList.Head.Value, nil
}

// GetTail ...
func (dList *DoublyLinkedList) GetTail() (int, error) {
	if dList.Tail == nil {
		return 0, errTailDoesNotExist
	}
	return dList.Tail.Value, nil
}

// Prepend (InsertAtFront) ...
func (dList *DoublyLinkedList) Prepend(value int) {
	node := Node{Value: value}

	if dList.Head == nil {
		dList.Head = &node
		dList.Tail = &node
		return
	}

	dList.Head.Previous = &node
	node.Next = dList.Head
	dList.Head = &node
}

// Append (InsertAtEnd) ...
func (dList *DoublyLinkedList) Append(value int) {
	node := Node{Value: value}

	if dList.Head == nil {
		dList.Head = &node
		dList.Tail = &node
		return
	}

	dList.Tail.Next = &node
	node.Previous = dList.Tail
	dList.Tail = &node
}

// InsertBetween ...
func (dList *DoublyLinkedList) InsertBetween(value, valuePrevious, valueNext int) error {
	Head := dList.Head

	if Head == nil {
		return errHeadDoesNotExist
	}

	node := Node{Value: value}

	for Head != nil && Head.Next != nil {
		if Head.Value == valuePrevious && (Head.Next).Value == valueNext {
			node.Previous, node.Next = Head, Head.Next
			(Head.Next).Previous = &node
			Head.Next = &node
			return nil
		}
		Head = Head.Next
	}

	return errRangeDoesNotExist
}

// RemoveHead ...
func (dList *DoublyLinkedList) RemoveHead() error {
	if dList.Head == nil {
		return errHeadDoesNotExist
	}

	if dList.Head == dList.Tail {
		dList.Head = nil
		dList.Tail = nil
		return nil
	}

	dList.Head = dList.Head.Next
	(dList.Head).Previous = nil

	return nil
}

// RemoveTail ...
func (dList *DoublyLinkedList) RemoveTail() error {
	if dList.Tail == nil {
		return errTailDoesNotExist
	}

	if dList.Tail == dList.Head {
		dList.Head = nil
		dList.Tail = nil
		return nil
	}

	dList.Tail = dList.Tail.Previous
	(dList.Tail).Next = nil

	return nil
}

// RemoveElement ...
func (dList *DoublyLinkedList) RemoveElement(value int) error {
	if dList.Head == nil {
		return errHeadDoesNotExist
	}

	node, err := dList.FindElement(value)
	if err != nil {
		return err
	}

	if node == dList.Head {
		return dList.RemoveHead()
	}

	if node == dList.Tail {
		return dList.RemoveTail()
	}

	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous

	return nil
}

// FindElement ...
func (dList *DoublyLinkedList) FindElement(value int) (*Node, error) {
	if dList.Head == nil {
		return nil, errHeadDoesNotExist
	}

	temp := dList.Head

	for temp != nil {
		if temp.Value == value {
			return temp, nil
		}
		temp = temp.Next
	}

	return nil, errElementDoesNotExist
}

// ClearList ...
func (dList *DoublyLinkedList) ClearList() {
	dList = &DoublyLinkedList{Head: nil, Tail: nil}
}

// traverseList ...
func traverseList(Head *Node) []int {
	if Head == nil {
		return []int{}
	}
	return append([]int{Head.Value}, traverseList(Head.Next)...)
}

// traverseListReverse ...
func traverseListReverse(Tail *Node) []int {
	if Tail == nil {
		return []int{}
	}
	return append([]int{Tail.Value}, traverseListReverse(Tail.Previous)...)
}

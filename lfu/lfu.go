package lfu

import (
	linkedList "github.com/el10savio/goLFUCache/linkedList"
)

// LFU ...
type LFU struct {
	Capacity      uint
	hashTable     map[int]int
	frequencyList linkedList.DoublyLinkedList
	accessList    linkedList.DoublyLinkedList
}

// InitLFU ...
// TODO: err check if capacity is 0
func InitLFU(capacity uint) *LFU {
	return &LFU{
		Capacity:      capacity,
		hashTable:     make(map[int]int),
		frequencyList: *linkedList.InitDoublyLinkedList(),
		accessList:    *linkedList.InitDoublyLinkedList(),
	}
}

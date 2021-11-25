package lfu

import (
	"errors"
	"log"

	linkedList "github.com/el10savio/goLFUCache/linkedList"
)

// LFU ...
type LFU struct {
	Capacity      uint
	hashTable     map[int]int
	accessList    map[*linkedList.Node][]CacheEntry
	frequencyList linkedList.DoublyLinkedList
}

// CacheEntry ...
type CacheEntry struct {
	Key   int
	Value int
}

// InitLFUCache ...
func InitLFUCache(capacity uint) (*LFU, error) {
	if capacity == 0 {
		return nil, errors.New("cannot create lfu cache with 0 capacity")
	}

	lfu := &LFU{
		Capacity:      capacity,
		hashTable:     make(map[int]int),
		accessList:    make(map[*linkedList.Node][]CacheEntry),
		frequencyList: *linkedList.InitDoublyLinkedList(),
	}

	return lfu, nil
}

// ClearLFUCache ...
func (lfu *LFU) ClearLFUCache() {
	lfu, _ = InitLFUCache(lfu.Capacity)
}

// NewEntry ...
func (lfu *LFU) NewEntry(key, value int) {
	// Add To HashTable
	lfu.hashTable[key] = value

	// Add "1" To frequencyList
	lfu.frequencyList.Append(1)

	// Add Entry To 1's AccessList
	frequencyNode, err := lfu.frequencyList.FindElement(1)
	if err != nil {
		log.Fatal(err)
	}

	lfu.accessList[frequencyNode] = append(lfu.accessList[frequencyNode], CacheEntry{key, value})
}

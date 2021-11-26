package lfu

import (
	"errors"

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
func (lfu *LFU) NewEntry(key, value int) error {
	if _, ok := lfu.hashTable[key]; ok {
		return errors.New("key already present in cache")
	}

	// Add To HashTable
	lfu.hashTable[key] = value

	// Add "1" To frequencyList
	lfu.frequencyList.Append(1)

	// Add Entry To 1's AccessList
	frequencyNode, err := lfu.frequencyList.FindElement(1)
	if err != nil {
		return err
	}

	lfu.accessList[frequencyNode] = append(lfu.accessList[frequencyNode], CacheEntry{key, value})

	return nil
}

// GetEntry ...
func (lfu *LFU) GetEntry(key int) (int, error) {
	if _, ok := lfu.hashTable[key]; !ok {
		return 0, errors.New("key is not present in cache")
	}

	// Update frequencyList
	// Find frequencyList Node containing key
	// Check if frequencyList's Next Node exists (==freq++)
	// If it does - move entry there
	// Else Create Node

	return lfu.hashTable[key], nil
}

// TODO: Error Handling
func (lfu *LFU) FindNode(key int) *linkedList.Node {
	head := lfu.frequencyList.Head

	for head != nil {
		for _, list := range lfu.accessList[head] {
			if key == list.Key {
				return head
			}
			head, _ = linkedList.GetNextElement(head)
		}
	}

	return nil
}

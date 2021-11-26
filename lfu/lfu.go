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
	node, listIndex := lfu.FindNode(key)

	// Check if frequencyList's Next Node exists
	next := node.Next

	if next == nil || next.Value != node.Value+1 {
		// Add node.Value+1 To frequencyList
		lfu.frequencyList.Append(node.Value + 1)

		// TODO: Handle case where there's another node &
		// the new node is to be added in between them
	}

	// TODO: Refactor To Seperate Function
	// Add Entry To AccessList
	frequencyNode, err := lfu.frequencyList.FindElement(node.Value + 1)
	if err != nil {
		return 0, err
	}

	lfu.accessList[frequencyNode] = append(lfu.accessList[frequencyNode], CacheEntry{key, lfu.hashTable[key]})

	// Delete Old frequencyNode's Entry
	lfu.accessList[node] = append(lfu.accessList[node][:listIndex], lfu.accessList[node][listIndex+1:]...)

	return lfu.hashTable[key], nil
}

// FindNode ...
func (lfu *LFU) FindNode(key int) (*linkedList.Node, int) {
	head := lfu.frequencyList.Head

	for head != nil {
		for index, list := range lfu.accessList[head] {
			if key == list.Key {
				return head, index
			}
			head = head.Next
		}
	}

	return nil, 0
}
